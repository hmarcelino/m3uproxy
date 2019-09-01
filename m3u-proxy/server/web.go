package server

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hmarcelino/m3u-proxy/config"
	"github.com/hmarcelino/m3u-proxy/server/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const Logo = ` 
___  ___ _____ _   _     ____________ _______   ____   __
|  \/  ||____ | | | |    | ___ \ ___ \  _  \ \ / /\ \ / /
| .  . |    / / | | |    | |_/ / |_/ / | | |\ V /  \ V / 
| |\/| |    \ \ | | |    |  __/|    /| | | |/   \   \ /  
| |  | |.___/ / |_| |    | |   | |\ \\ \_/ / /^\ \  | |  
\_|  |_/\____/ \___/     \_|   \_| \_|\___/\/   \/  \_/

is accepting requests in port :%d
* http://127.0.0.1:%d
* http://%s:%d

`

func Start(config *config.Config) {
	muxRouter := mux.NewRouter()

	register(muxRouter, config, routes.RootRouter)
	register(muxRouter, config, routes.PingRouter)
	register(muxRouter, config, routes.ChannelListRouter)
	register(muxRouter, config, routes.ChannelRoute)
	register(muxRouter, config, routes.ChannelInfoRoute)

	fmt.Printf(
		Logo,
		config.Server.Port,
		config.Server.Port,
		config.Server.Hostname,
		config.Server.Port)

	server := &http.Server{Addr: fmt.Sprintf(":%d", config.Server.Port), Handler: muxRouter}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			//log.Fatalf("Error starting server: %v", err)
		}
	}()

	_, err := routes.LoadList(config)
	if routes.LoadList(config); err != nil {
		log.Fatalf(err.Msg+" %v", err.Error)
	}

	log.Println("List loaded successfully")

	// Setting up signal capturing
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Waiting for SIGINT (pkill -2)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		// ignoring error
	}
}

func register(mux *mux.Router, config *config.Config, route func(config *config.Config) (string, func(w http.ResponseWriter, r *http.Request))) {
	path, handler := route(config)
	mux.HandleFunc(path, handler)
}
