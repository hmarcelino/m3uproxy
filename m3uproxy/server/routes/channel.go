package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hmarcelino/m3uproxy/config"
	"github.com/hmarcelino/m3uproxy/db"
	"github.com/hmarcelino/m3uproxy/server/webutils"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const QueryParamLocation = "location"
const HeaderChannelId = "X-ChannelId"
const HeaderRange = "Range"

func ChannelRoute(config *config.Config) (string, func(w http.ResponseWriter, r *http.Request)) {
	var responseModifier = GetResponseModifier(config)

	return "/channels/{id}", func(w http.ResponseWriter, r *http.Request) {

		var channelAddr *db.Channel
		var err error

		vars := mux.Vars(r)
		channelId := vars["id"]

		// Decide if we want to lookup from the database
		// or use the url provided in the request query parameter
		overrideChannelAddr := r.URL.Query().Get(QueryParamLocation)
		if overrideChannelAddr != "" {
			newUrl, err := url.Parse(overrideChannelAddr)
			if err != nil {
				webutils.BadRequest("Invalid channel address override: "+overrideChannelAddr, err, w)
				return
			}

			channelAddr = &db.Channel{Id: channelId, Source: newUrl}

		} else {

			// Fallback to lookup in the database
			// if no override channel address is provided
			channelAddr, err = db.LookupChannel(channelId)
			if err != nil {
				webutils.NotFound(w)
				return
			}
		}

		dump, err := httputil.DumpRequest(r, false)
		log.Printf("%q\r\nRemoteAddr: %s", dump, r.RemoteAddr)

		request := http.Request{URL: channelAddr.Source}
		request.Header = map[string][]string{}
		request.Header.Add(HeaderChannelId, channelId)

		rangeValue := r.Header.Get(HeaderRange)
		if rangeValue != "" {
			request.Header.Add(HeaderRange, rangeValue)
		}

		proxy := newProxy(channelAddr)
		if overrideChannelAddr == "" {
			proxy.ModifyResponse = responseModifier
		}

		log.Printf("Proxying request for channel %s %s redirect=%t",
			channelId,
			channelAddr.Source.String(),
			overrideChannelAddr != "")

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
			newReq, _ := url.Parse(GetChannelUrl(config, channelId))
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
func GetChannelUrl(config *config.Config, id string) string {
	return fmt.Sprintf(
		"http://%s:%d/channels/%s",
		config.Server.Hostname,
		config.Server.Port,
		id,
	)
}
