package main

import (
	"fmt"
	"github.com/linda-lai/golang-ticket-viewer/internal/api"
	"github.com/linda-lai/golang-ticket-viewer/internal/config"
	"github.com/linda-lai/golang-ticket-viewer/internal/cli"
)


func main() {
	// Initialise config
	credentials, subdomain := config.Auth()
	endpoint := api.GetUrl(subdomain)

	// Show Ticket By ID
	//ticketID := "20"
	//ticket := cli.ShowTicket(credentials, endpoint, ticketID)
	//
	//barrier := "============================================"
	//fmt.Println(barrier)
	//fmt.Println(ticket.Ticket.ID)
	//fmt.Println(barrier)
	//fmt.Println(ticket.Ticket.Subject)
	//fmt.Println(barrier)
	//fmt.Println(ticket.Ticket.Description)
	//fmt.Println(barrier)

	// List All Tickets
	//fmt.Println(cli.ListAllTickets(credentials, endpoint))
	barrier := "|------------||---------------------------------------------------------||--------------------------------|"
	values := map[string]string{"ticketID": "TICKET ID", "subject": "SUBJECT", "createdAt": "CREATED AT"}
	fmt.Println("\n" + barrier)
	fmt.Printf("| %-10v || %-55v || %-31v|", values["ticketID"], values["subject"], values["createdAt"])
	fmt.Println("\n" + barrier)

	var tickets = cli.ListAllTickets(credentials, endpoint)

	for i := range(tickets.Tickets) {
		fmt.Printf("| %-10v || %-55v || %-30v |\n", tickets.Tickets[i].ID, tickets.Tickets[i].Subject, tickets.Tickets[i].CreatedAt)
	}

	fmt.Println(barrier)

	fmt.Printf("| TOTAL: %v ||\n", tickets.Count)
}
