package zendesk

import (
	"encoding/json"
	"fmt"

	"github.com/linda-lai/golang-ticket-viewer/internal/api"
	"github.com/linda-lai/golang-ticket-viewer/internal/config"

)

// type ZendeskTickets struct {
// 	count int
// 	tickets []ZendeskTicketByID
// }

type ZendeskTicket struct {
	Ticket struct {
		ID int `json:"id"`
		Subject string `json:"subject"`
		Description string `json:"description"`
	} `json:"ticket"`
}

func getTicket() {
	// Get Ticket By ID
	// fmt.Println(api.GetTicketByID(config.Username, config.Password, config.Subdomain, "20"))

	var ticket ZendeskTicket

	data := api.GetTicketByID(config.Username, config.Password, config.Subdomain, "1")
	fmt.Println(data)

	err := json.Unmarshal([]byte(data), &ticket)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("+++++++++++++++++++++++++++++")

	fmt.Printf("%+v\n", ticket)

	fmt.Println("----------------------------")

	// // List All Tickets
	// fmt.Println(api.GetTickets(config.Username, config.Password, config.Subdomain))

	// fmt.Println("hello")
}