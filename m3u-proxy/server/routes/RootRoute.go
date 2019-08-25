package routes

import (
	"fmt"
	"github.com/hmarcelino/m3u-proxy/config"
	"log"
	"net/http"
)

func RootRouter(config *config.Config) (string, func(w http.ResponseWriter, r *http.Request)) {
	return "/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Welcome to m3u proxy")
		if err != nil {
			log.Printf("Error writing to output welcome message: %v", err)
		}
	}
}
