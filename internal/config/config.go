package config

import (
	"os"
)

type ZendeskCredentials struct {
	Username string
	Password string
	Subdomain string
}

func New() *ZendeskCredentials {
	return &ZendeskCredentials{
		Username:  getEnv("ZD_USERNAME", ""),
		Password:  getEnv("ZD_PASSWORD", ""),
		Subdomain: getEnv("ZD_SUBDOMAIN", ""),
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key, defaultEnv string) string {
	if env, exists := os.LookupEnv(key); exists {
		return env
	}

	return defaultEnv
}