package routes

import (
	"github.com/hmarcelino/m3u-proxy/config"
	"github.com/hmarcelino/m3u-proxy/server/webutils"
	"net/http"
)

var resp = []byte("Welcome to m3u proxy")

func RootRouter(config *config.Config) (string, func(w http.ResponseWriter, r *http.Request)) {
	return "/", func(w http.ResponseWriter, r *http.Request) {
		webutils.Success(resp, w)
	}
}
