package config

import (
	"encoding/base64"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Zendesk credentials for Basic Auth
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

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func Auth() (string, string) {
	config := New()
	subdomain := config.Subdomain
	auth := basicAuth(config.Username, config.Password)
	return auth, subdomain
}

