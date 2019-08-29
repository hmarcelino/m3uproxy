package main

import (
	"flag"
	"github.com/hmarcelino/m3u-proxy/config"
	"github.com/hmarcelino/m3u-proxy/server"
)

func main() {
	var ymlConfigFile string
	flag.StringVar(&ymlConfigFile, "file", "", "Configuration file")
	flag.Parse()

	var m3uServerConfig *config.Config

	if ymlConfigFile != "" {
		m3uServerConfig = config.LoadYml(ymlConfigFile)

	} else {
		m3uServerConfig = config.LoadEnv()
	}

	config.Validate(m3uServerConfig)
	server.Start(m3uServerConfig)
}
