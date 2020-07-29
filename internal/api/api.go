package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Set Client to an instance of HTTPClient interface
var (
	Client HTTPClient
)

// HTTPClient interface implements the Do function. Takes a pointer to
// an http.Request and returns either a pointer to a response or error.
// This is the exact API of the existing http.Client's Do function.
type HTTPClient interface {
	 Do(req *http.Request) (*http.Response, error)
}

func init() {
	Client = &http.Client{}
}

// newZendeskClient returns an unmarshalled JSON string with Zendesk ticket
// or tickets data from the Ticket API
func newZendeskClient(url, auth string) (*http.Response, error) {
	// Client is a pointer to an instance of http.Client struct,
	// so it can be shared and reused.
	//client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", "Basic " + auth)
	//req.SetBasicAuth(username, password)


	if err != nil {
		log.Fatal(err)
	}

	return Client.Do(req)
}

// GetUrl constructs a generic Ticket endpoint with the user subdomain
func GetUrl(subdomain string) string {
	return fmt.Sprintf("https://%s.zendesk.com/api/v2/tickets", subdomain)
}

// GetTicketByID constructs the endpoint to get a Zendesk ticket by ID
func GetTicketByID(auth, url, ticketID string) (string, error) {
	ticketUrl :=  fmt.Sprintf("%s/%s.json", url, ticketID)
	resp, err := newZendeskClient(ticketUrl, auth)

	defer resp.Body.Close()

	// The ioutil package implements I/O utility functions, so we can
	// read and return the entire response body with io.Reader interface
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	data := string(body)

	return data, nil
}

// ListTickets constructs the endpoint to get all Zendesk tickets,
// passing it to a new Zendesk client
func ListTickets(auth, url string) (string, error) {
	pagination := "25"
	ticketsUrl := fmt.Sprintf("%s.json?per_page=%s", url, pagination)

	resp, err := newZendeskClient(ticketsUrl, auth)

	defer resp.Body.Close()

	// The ioutil package implements I/O utility functions, so we can
	// read and return the entire response body with io.Reader interface
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	data := string(body)

	return data, nil
}
