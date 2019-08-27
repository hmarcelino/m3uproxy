package config

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

func Validate(config *Config) {
	err := joinValidationErrors(
		validateServerPort(config),
		validateM3uUrl(config),
	)

	if err != nil {
		log.Fatalf("Invalid configuration: %v\n", err)
	}
}

func validateServerPort(config *Config) error {
	if config.Server.Port == 0 {
		return newValidationError("Invalid port for server. You must specify port number bigger than 1024")
	}
	return nil
}

func validateM3uUrl(config *Config) error {
	m3uUrl := strings.TrimSpace(config.M3u.Url)
	if m3uUrl == "" {
		return newValidationError("You must specify a m3u URL address")
	}
	_, err := url.Parse(m3uUrl)
	if err != nil {
		return newValidationError(fmt.Sprintf("Error parsing m3u Url %s \n", m3uUrl))
	}
	return nil
}

func newValidationError(s string) error {
	return fmt.Errorf("* %s \n", s)
}

func joinValidationErrors(errors ...error) error {
	var joinErrors = ""

	for _, e := range errors {
		if e != nil {
			joinErrors += e.Error()
		}
	}

	if joinErrors != "" {
		return fmt.Errorf(joinErrors)
	}

	return nil
}
