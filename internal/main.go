package main

import (
	"fmt"

	"github.com/linda-lai/golang-ticket-viewer/internal/api"
	"github.com/linda-lai/golang-ticket-viewer/internal/cli"
	"github.com/linda-lai/golang-ticket-viewer/internal/config"
	"github.com/linda-lai/golang-ticket-viewer/internal/zendesk"
)

func main() {
	// Initialise config
	credentials, subdomain := config.Auth()
	endpoint := api.GetUrl(subdomain)

	// Show Ticket By ID
	ticketID := "20"
	ticket := zendesk.ShowTicket(credentials, endpoint, ticketID)
	fmt.Println(cli.ShowTicket(ticket))

	// List All Tickets
	var tickets = zendesk.ListTicketsData(credentials, endpoint)
	fmt.Println(cli.ListAllTickets(tickets))

	//fmt.Print("Enter text: \n")
	//var input string
	//_, err := fmt.Scanln(&input)
	//if err != nil {
	//	fmt.Fprintln(os.Stderr, err)
	//	return
	//}
	//fmt.Println(input)
}
