package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/hmarcelino/m3uproxy/config"
	"github.com/hmarcelino/m3uproxy/db"
	"github.com/hmarcelino/m3uproxy/server/webutils"
	"net/http"
)

func ChannelInfoRoute(config *config.Config) (string, func(w http.ResponseWriter, r *http.Request)) {
	return "/channels/info/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		channelId := vars["id"]

		channel, err := db.GetChannel(channelId)
		if err != nil {
			webutils.NotFound(w)
		}

		jsonChannel, err := json.Marshal(channel)
		if err != nil {
			webutils.InternalServerError("Error building jsonChannel from a db.Channel", err, w)
		}

		w.Header().Set("Content-Type", "application/json")
		webutils.Success(jsonChannel, w)
	}
}
