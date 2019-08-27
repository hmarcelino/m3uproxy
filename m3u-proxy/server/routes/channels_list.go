package routes

import (
	"fmt"
	"github.com/hmarcelino/m3u-proxy/config"
	"github.com/hmarcelino/m3u-proxy/db"
	"github.com/hmarcelino/m3u-proxy/server/webutils"
	"io/ioutil"
	"net/http"
	"strings"
)

func ChannelListRouter(config *config.Config) (string, func(w http.ResponseWriter, r *http.Request)) {
	return "/channels", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get(config.M3u.Url)
		if err != nil {
			webutils.BadGateway("Error loading channels list", err, w)
			return
		}

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			webutils.InternalServerError("Error loading body from response", err, w)
			return
		}

		err = resp.Body.Close()
		if err != nil {
			webutils.InternalServerError("Error closing body response", err, w)
			return
		}

		if resp.StatusCode == 200 {
			db.Reset()

			b, err = modifyResponse(config, string(b))
			if err != nil {
				webutils.InternalServerError("Error modifying response", err, w)
				return
			}
		}

		webutils.Success(b, w)
	}
}

func modifyResponse(config *config.Config, payload string) ([]byte, error) {
	lines := strings.Split(string(payload), "\r\n")

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		// Line holding metadata. Something like
		// #EXTINF:-1 tvg-id="ABC HD" tvg-name="ABC FHD" ...
		if !strings.HasPrefix(line, "http://") {
			continue
		}

		// line is a channel address.
		// Override channel address with proxyHost address
		channel, err := db.RegisterChannel(line)
		if err != nil {
			return nil, fmt.Errorf("error registering m3u url. %v", err)
		}

		lines[i] = NewChannelProxy(config, channel.Id)
	}

	return []byte(strings.Join(lines, "\n")), nil
}
