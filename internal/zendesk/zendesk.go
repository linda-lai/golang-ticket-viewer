package zendesk

import (
	"encoding/json"
	"fmt"
	"github.com/linda-lai/golang-ticket-viewer/internal/api"
)

type TicketsPayload struct {
	Tickets []TicketFields`json:"tickets",omitempty`
	NextPage string `json:"next_page",omitempty`
	PreviousPage string `json:"previous_page",omitempty`
	Count int `json:count`
}

type TicketPayload struct {
	Ticket TicketFields `json:"ticket",omitempty`
}

type TicketFields struct {
	ID int `json:"id,omitempty"`
	Subject string `json:"subject,omitempty"`
	Description string `json:"description,omitempty"`
}

func UnmarshalZendeskTicket(credentials, endpoint, ticketID string) string {
	var ticket TicketPayload
	data := api.GetTicketByID(credentials, endpoint, ticketID)

	err := json.Unmarshal([]byte(data), &ticket)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("+++++++++++++++++++++++++++++")
	fmt.Printf("%+v\n", ticket)
	fmt.Println("----------------------------")

	return data
}

func UnmarshalZendeskTickets(credentials, endpoint string) string {
	var tickets TicketsPayload
	data := api.GetTickets(credentials, endpoint)
	//fmt.Println(data)

	err := json.Unmarshal([]byte(data), &tickets)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("+++++++++++++++++++++++++++++")
	fmt.Printf("%+v\n", tickets)
	fmt.Println("----------------------------")

	return data

}