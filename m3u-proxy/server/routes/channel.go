package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hmarcelino/m3u-proxy/config"
	"github.com/hmarcelino/m3u-proxy/db"
	"github.com/hmarcelino/m3u-proxy/server/webutils"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const QueryParamLocation = "location"
const HeaderChannelId = "X-ChannelId"
const HeaderRange = "Range"

func ChannelRoute(config *config.Config) (string, func(w http.ResponseWriter, r *http.Request)) {
	return "/channels/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		channelId := vars["id"]

		channelAddr, err := db.LookupChannel(channelId)

		query := r.URL.Query()
		newChannelAddr := query.Get(QueryParamLocation)
		rangeValue := r.Header.Get(HeaderRange)

		if err != nil {
			webutils.NotFound(w)
			return
		}

		dump, err := httputil.DumpRequest(r, false)
		log.Printf("%q\r\nRemoteAddr: %s", dump, r.RemoteAddr)

		if newChannelAddr != "" {
			newUrl, err := url.Parse(newChannelAddr)
			if err != nil {
				webutils.BadRequest("Invalid channel address override: "+newChannelAddr, err, w)
				return
			}

			channelAddr = &db.Channel{Id: channelId, Source: newUrl}
		}

		request := http.Request{URL: channelAddr.Source}
		request.Header = map[string][]string{}
		request.Header.Add(HeaderChannelId, channelId)

		if rangeValue != "" {
			request.Header.Add(HeaderRange, rangeValue)
		}

		proxy := newProxy(channelAddr)
		if newChannelAddr == "" {
			proxy.ModifyResponse = GetResponseModifier(config)
		}

		log.Printf("Proxying request for channel %s %s redirect=%t",
			channelId,
			channelAddr.Source.String(),
			newChannelAddr != "")

		proxy.ServeHTTP(w, &request)
	}
}

func newProxy(channel *db.Channel) *httputil.ReverseProxy {
	addr := channel.Source
	uHost, _ := url.Parse(addr.Scheme + "://" + addr.Host)
	return httputil.NewSingleHostReverseProxy(uHost)
}

func GetResponseModifier(config *config.Config) func(resp *http.Response) error {
	return func(resp *http.Response) error {
		isRedirect := resp.StatusCode == http.StatusFound ||
			resp.StatusCode == http.StatusSeeOther ||
			resp.StatusCode == http.StatusTemporaryRedirect

		channelId := resp.Request.Header.Get(HeaderChannelId)

		if isRedirect && channelId != "" {
			newReq, _ := url.Parse(GetChanneUrl(config, channelId))
			query := newReq.Query()
			query.Set(QueryParamLocation, resp.Header.Get("Location"))

			newReq.RawQuery = query.Encode()
			resp.Header.Set("Location", newReq.String())
		}

		return nil
	}
}

// The return should match the previous route pattern.
// Http://host:port/channels/channelId
func GetChanneUrl(config *config.Config, id string) string {
	return fmt.Sprintf(
		"http://%s:%d/channels/%s",
		config.Server.Hostname,
		config.Server.Port,
		id,
	)
}
