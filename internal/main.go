package main

import (
	"fmt"
	"github.com/linda-lai/golang-ticket-viewer/internal/zendesk"

	"github.com/linda-lai/golang-ticket-viewer/internal/api"
	"github.com/linda-lai/golang-ticket-viewer/internal/config"
)


func main() {
	// Initialise config
	credentials, subdomain := config.Auth()
	endpoint := api.GetUrl(subdomain)

	// Show Ticket By ID
	ticketID := "20"
	fmt.Println(zendesk.UnmarshalZendeskTicket(credentials, endpoint, ticketID))

	// List All Tickets
	fmt.Println(zendesk.UnmarshalZendeskTickets(credentials, endpoint))

}
