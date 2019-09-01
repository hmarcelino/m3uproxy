package config

import (
	"log"
	"os"
	"strconv"
)

const (
	M3uProxyPort     = "M3U_PROXY_PORT"
	M3uProxyHostname = "M3U_PROXY_HOSTNAME"
	M3uProxyM3uUrl   = "M3U_PROXY_CHANNELS_URL"
)

func LoadEnv() *Config {
	var config = &Config{}
	config.Server.Port = 9090
	config.Server.Hostname = "localhost"

	port := os.Getenv(M3uProxyPort)
	if port != "" {
		envPort, err := strconv.ParseInt(port, 10, 64)
		if err != nil {
			log.Fatalf("Error parsing server port number: %s", port)
		}

		config.Server.Port = uint16(envPort)
	}

	envHostname := os.Getenv(M3uProxyHostname)
	if envHostname != "" {
		config.Server.Hostname = envHostname
	}

	config.M3u.Url = os.Getenv(M3uProxyM3uUrl)
	return config
}
