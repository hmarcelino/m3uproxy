package main

import (
	"flag"
	"fmt"
	"github.com/hmarcelino/m3u-proxy/config"
	"log"
	"strings"
)

const logo = ` 
___  ___ _____ _   _     ____________ _______   ____   __
|  \/  ||____ | | | |    | ___ \ ___ \  _  \ \ / /\ \ / /
| .  . |    / / | | |    | |_/ / |_/ / | | |\ V /  \ V / 
| |\/| |    \ \ | | |    |  __/|    /| | | |/   \   \ /  
| |  | |.___/ / |_| |    | |   | |\ \\ \_/ / /^\ \  | |  
\_|  |_/\____/ \___/     \_|   \_| \_|\___/\/   \/  \_/

is accepting requests in port :%d

===========================================
http://%s:%d'

`

func main() {
	var ymlConfigFile string
	flag.StringVar(&ymlConfigFile, "file", "", "Configuration file")
	flag.Parse()

	ymlConfigFile = strings.TrimSpace(ymlConfigFile)
	if len(ymlConfigFile) == 0 {
		log.Fatalf("No configuration file provided")
	}

	config, err := config.LoadYmlFile(ymlConfigFile)
	if err != nil {
		log.Fatalf("Error reading configuration file: %v", err)
	}

	fmt.Printf(logo, config.Server.Port, config.Server.Hostname, config.Server.Port)
}
