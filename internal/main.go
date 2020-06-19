package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"os"

	"github.com/joho/godotenv"

	"github.com/linda-lai/golang-ticket-viewer/internal/config"
)


func newZendeskClient(url, username, password string) string {
	// Client is a pointer to an instance of http.Client struct,
	// so it can be shared and reused.
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	//req.Header.Add("Authorization", "Basic "+basicAuth(username, password))
	req.SetBasicAuth(username, password)

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	// Close response body to prevent resource leaks
	defer resp.Body.Close()

	// The ioutil package implements I/O utility functions, so we can
	// read and return the entire response body with io.Reader interface
	body, err := ioutil.ReadAll(resp.Body)
	data := string(body)

	return data
}

func getTicketById(username, password, subdomain, ticketId string) string {
	url :=  fmt.Sprintf("https://%s.zendesk.com/api/v2/tickets/%s.json", subdomain, ticketId)

	return newZendeskClient(url, username, password)
}

func getTickets(username, password, subdomain string) string {
	pagination := "2"
	url := fmt.Sprintf("https://%s.zendesk.com/api/v2/tickets.json?per_page=%s", subdomain, pagination)

	return newZendeskClient(url, username, password)
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	config := config.New()

	// Get Ticket By ID
	fmt.Println(getTicketById(config.Username, config.Password, config.Subdomain, "1"))

	fmt.Println("----------------------------")

	// List All Tickets
	fmt.Println(getTickets(config.Username, config.Password, config.Subdomain))

}
