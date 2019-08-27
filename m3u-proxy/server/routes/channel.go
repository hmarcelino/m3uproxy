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

const QUERY_PARAM_LOCATION = "location"
const CHANNEL_ID_HEADER = "X-ChannelId"

func ChannelRoute(config *config.Config) (string, func(w http.ResponseWriter, r *http.Request)) {
	return "/channels/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		channelId := vars["id"]

		channelAddr, err := db.LookupChannel(channelId)
		if err != nil {
			webutils.NotFound(w)
			return
		}

		dump, err := httputil.DumpRequest(r, false)
		log.Printf("%q\r\nRemoteAddr: %s", dump, r.RemoteAddr)

		// Channel location override
		location := r.URL.Query().Get(QUERY_PARAM_LOCATION)
		if location != "" {
			newUrl, err := url.Parse(location)
			if err != nil {
				webutils.BadGateway("Invalid location address", err, w)
				return
			}

			log.Printf("Got new location for channel: %s, location: %s\n", channelId, location)
			newLocationChannelAddr := &db.Channel{Id: channelId, Source: newUrl}
			newProxy(config, newLocationChannelAddr).ServeHTTP(w, &http.Request{URL: newLocationChannelAddr.Source})
			return
		}

		// fallBack to channel address
		request := http.Request{URL: channelAddr.Source}
		request.Header = map[string][]string{}
		request.Header.Add(CHANNEL_ID_HEADER, channelId)
		newProxy(config, channelAddr).ServeHTTP(w, &request)
	}
}

func newProxy(config *config.Config, channel *db.Channel) *httputil.ReverseProxy {
	addr := channel.Source
	uHost, _ := url.Parse(addr.Scheme + "://" + addr.Host)

	proxy := httputil.NewSingleHostReverseProxy(uHost)
	proxy.ModifyResponse = func(resp *http.Response) error {

		isRedirect := resp.StatusCode == http.StatusFound ||
			resp.StatusCode == http.StatusSeeOther ||
			resp.StatusCode == http.StatusTemporaryRedirect

		channelId := resp.Request.Header.Get(CHANNEL_ID_HEADER)

		if isRedirect && channelId != "" {
			newReq, _ := url.Parse(NewChannelProxy(config, channelId))
			query := newReq.Query()
			query.Set(QUERY_PARAM_LOCATION, resp.Header.Get("Location"))

			newReq.RawQuery = query.Encode()
			resp.Header.Set("Location", newReq.String())
		}

		return nil
	}

	return proxy
}

// The return should match the previous route pattern.
// Http://host:port/channels/channelId
func NewChannelProxy(config *config.Config, id string) string {
	return fmt.Sprintf(
		"http://%s:%d/channels/%s",
		config.Server.Hostname,
		config.Server.Port,
		id,
	)
}
