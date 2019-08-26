package routes

import (
	"github.com/hmarcelino/m3u-proxy/config"
	"github.com/hmarcelino/m3u-proxy/server/webutils"
	"net/http"
)

var bytes = []byte("pong")

func PingRouter(config *config.Config) (string, func(w http.ResponseWriter, r *http.Request)) {
	return "/ping", func(w http.ResponseWriter, r *http.Request) {
		webutils.Success(bytes, w)
	}
}
