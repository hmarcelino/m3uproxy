package main

import (
	"flag"
	"github.com/hmarcelino/m3uproxy/config"
	"github.com/hmarcelino/m3uproxy/server"
)

func main() {
	var m3uServerConfig *config.Config

	var ymlConfigFile string
	flag.StringVar(&ymlConfigFile, "file", "", "Configuration file")
	flag.Parse()

	if ymlConfigFile != "" {
		m3uServerConfig = config.LoadYml(ymlConfigFile)

	} else {
		m3uServerConfig = config.LoadEnv()
	}

	config.Validate(m3uServerConfig)
	server.Start(m3uServerConfig)
}
