package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func basicAuth(username, password string) string {
	return base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
}

func getTickets(username, password, subdomain string) string {
	client := &http.Client{}
	pagination := "2"
	url := fmt.Sprintf("https://%s.zendesk.com/api/v2/tickets.json?per_page=%s", subdomain, pagination)
	req, err := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", "Basic "+basicAuth(username, password))
	// req.SetBasicAuth(username, password)

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	tickets := string(body)

	return tickets
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	username := os.Getenv("ZD_USERNAME")
	password := os.Getenv("ZD_PASSWORD")
	subdomain := os.Getenv("ZD_SUBDOMAIN")
	fmt.Println(basicAuth(username, password))
	fmt.Println(getTickets(username, password, subdomain))
}
