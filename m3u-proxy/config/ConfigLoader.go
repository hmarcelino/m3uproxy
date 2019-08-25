package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Config struct {
	Server struct {
		Port     int
		Hostname string
	}

	M3u struct {
		Url string
	}
}

func LoadYml(ymlConfigFile string) *Config {
	ymlConfigFile = strings.TrimSpace(ymlConfigFile)
	if len(ymlConfigFile) == 0 {
		log.Fatalf("No configuration file provided")
	}

	m3uServerConfig, err := readYmlFile(ymlConfigFile)
	if err != nil {
		log.Fatalf("Error reading configuration file: %v", err)
	}
	return m3uServerConfig
}

func readYmlFile(configFilePath string) (*Config, error) {
	var config = Config{}
	config.Server.Port = 9090
	config.Server.Hostname = "localhost"

	yamlBytes, err := readFile(configFilePath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlBytes, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func readFile(configFilePath string) ([]byte, error) {
	var yamlBytes []byte

	file, err := os.Open(configFilePath)
	if err != nil {
		return yamlBytes, err
	}

	defer file.Close()

	yamlConfig, err := ioutil.ReadAll(file)
	if err != nil {
		return yamlBytes, err
	}

	return yamlConfig, nil
}
