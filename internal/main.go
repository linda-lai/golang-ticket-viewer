package main

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


func main() {

	// Get Ticket By ID
	// fmt.Println(api.GetTicketByID(config.Username, config.Password, config.Subdomain, "20"))

	var ticket ZendeskTicket

	credentials, subdomain := config.Auth()

	endpoint := api.GetUrl(subdomain)
	data := api.GetTicketByID(credentials, endpoint, "1")
	fmt.Println(data)

	err := json.Unmarshal([]byte(data), &ticket)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("+++++++++++++++++++++++++++++")

	fmt.Printf("%+v\n", ticket)

	fmt.Println("----------------------------")

	// // List All Tickets
	fmt.Println(api.GetTickets(credentials, endpoint))

	// fmt.Println("hello")

}
