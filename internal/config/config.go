package config

import (
	"encoding/base64"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	ZendeskUsernameEnv  = "ZD_USERNAME"
	ZendeskPasswordEnv  = "ZD_PASSWORD"
	ZendeskSubdomainEnv = "ZD_SUBDOMAIN"
	defaultEnv          = ""
)

// ZendeskCredentials struct for Basic Auth
type ZendeskCredentials struct {
	Username  string
	Password  string
	Subdomain string
	BasicAuth string
}

// New() returns a pointer to an instance of ZendeskCredentials,
// set to ENV variables or empty string
func New() *ZendeskCredentials {
	username := getEnv(ZendeskUsernameEnv, defaultEnv)
	password := getEnv(ZendeskPasswordEnv, defaultEnv)
	subdomain := getEnv(ZendeskSubdomainEnv, defaultEnv)
	auth := basicAuth(username, password)

	return &ZendeskCredentials{
		Username:  username,
		Password:  password,
		Subdomain: subdomain,
		BasicAuth: auth,
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key, defaultEnv string) string {
	if env, exists := os.LookupEnv(key); exists {
		return env
	}
	return defaultEnv
}

// Load will read your env file(s) and load them into ENV for this process.
// Call this function as close as possible to the start of your program (ideally in main)
// If you call Load without any args it will default to loading .env in the current path
// You can otherwise tell it which files to load (there can be more than one)
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// basicAuth encodes Zendesk username and passwords
// to base64 as a string
func basicAuth(username, password string) string {
	auth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	return auth
}

//// Auth intialises a new config instance and returns an encoded
//// Basic Auth and a Zendesk subdomain string
//func Auth() *ZendeskCredentials {
//	config := New()
//	return config
//}
//
