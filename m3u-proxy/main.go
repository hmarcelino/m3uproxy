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

	m3uServerConfig := config.LoadYml(ymlConfigFile)
	config.Validate(m3uServerConfig)
	server.Start(m3uServerConfig)
}
