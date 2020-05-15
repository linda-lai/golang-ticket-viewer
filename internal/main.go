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

func getTicketById(username, password, subdomain, ticketId string) string {
	client := &http.Client{}
	url := fmt.Sprintf("https://%s.zendesk.com/api/v2/tickets/%s.json", subdomain, ticketId)
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(username, password)

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	ticket := string(body)

	return ticket
}

func getTickets(username, password, subdomain string) string {
	// Client is a pointer to an instance of http.Client struct,
	// so it can be shared and reused.
	client := &http.Client{}
	pagination := "25"
	url := fmt.Sprintf("https://%s.zendesk.com/api/v2/tickets.json?per_page=%s", subdomain, pagination)
	req, err := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", "Basic "+basicAuth(username, password))

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	// Close response body to prevent resource leaks
	defer resp.Body.Close()

	// The ioutil package implements I/O utility functions, so we can
	// read and return the entire response body with io.Reader interface
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
	fmt.Println(getTicketById(username, password, subdomain, "1"))
}
