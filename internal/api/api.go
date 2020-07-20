package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// newZendeskClient returns an unmarshalled JSON string with Zendesk ticket
// or tickets data from the Ticket API
func newZendeskClient(url, auth string) string {
	// Client is a pointer to an instance of http.Client struct,
	// so it can be shared and reused.
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", "Basic " + auth)
	//req.SetBasicAuth(username, password)

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

// GetUrl constructs a generic Ticket endpoint with the user subdomain
func GetUrl(subdomain string) string {
	return fmt.Sprintf("https://%s.zendesk.com/api/v2/tickets", subdomain)
}

// GetTicketByID constructs the endpoint to get a Zendesk ticket by ID
func GetTicketByID(auth, url, ticketID string) string {
	ticketUrl :=  fmt.Sprintf("%s/%s.json", url, ticketID)

	return newZendeskClient(ticketUrl, auth)
}

// ListTickets constructs the endpoint to get all Zendesk tickets,
// passing it to a new Zendesk client
func ListTickets(auth, url string) string {
	pagination := "25"
	ticketsUrl := fmt.Sprintf("%s.json?per_page=%s", url, pagination)

	return newZendeskClient(ticketsUrl, auth)
}
