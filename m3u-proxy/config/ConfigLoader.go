package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	Server struct {
		Port     int
		Hostname string
	}
	M3u struct {
		OriginUrl string
	}
}

func LoadYmlFile(configFilePath string) (Config, error) {
	var config = Config{}
	config.Server.Port = 9090
	config.Server.Hostname = "localhost"

	yamlBytes, err := readFile(configFilePath)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal([]byte(yamlBytes), &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

// ++++++++++++++++++++++++++++
// Helper methods
// ++++++++++++++++++++++++++++

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
