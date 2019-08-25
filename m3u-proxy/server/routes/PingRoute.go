package routes

import (
	"fmt"
	"github.com/hmarcelino/m3u-proxy/config"
	"log"
	"net/http"
)

func PingRouter(config *config.Config) (string, func(w http.ResponseWriter, r *http.Request)) {
	return "/ping", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "pong")
		if err != nil {
			log.Printf("Error writing to output pong response: %v", err)
		}
	}
}
