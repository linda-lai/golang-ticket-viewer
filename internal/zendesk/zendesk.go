package zendesk

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/linda-lai/golang-ticket-viewer/internal/api"
)

// TicketsPayload contains all tickets data returned from Zendesk API
type TicketsPayload struct {
	Tickets []TicketFields`json:"tickets,omitempty"`
	NextPage string `json:"next_page,omitempty"`
	PreviousPage string `json:"previous_page,omitempty"`
	Count int `json:"count,omitempty"`
}

// TicketPayload contains a single ticket's data returned from Zendesk API
type TicketPayload struct {
	Ticket TicketFields `json:"ticket,omitempty"`
}

// TicketFields contains ticket fields and details nested in the
// TicketPayload or TicketsPayload
type TicketFields struct {
	ID int `json:"id,omitempty"`
	Subject string `json:"subject,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// UnmarshalZendeskTicket parses the response and maps the returned fields
// to the TicketPayload
// ShowTicket returns an instance of the TicketPayload data
func UnmarshalZendeskTicket(credentials, endpoint, ticketID string) TicketPayload {
	var ticket TicketPayload
	data, _ := api.GetTicketByID(credentials, endpoint, ticketID)

	err := json.Unmarshal([]byte(data), &ticket)

	if err != nil {
		fmt.Println(err)
	}

	//fmt.Printf("%+v\n", ticket)

	return ticket
}

// UnmarshalZendeskTickets parses the response and maps the returned fields
// to the TicketsPayload
// ListTicketsData returns an instance of the TicketsPayload data
func UnmarshalZendeskTickets(credentials, endpoint string) TicketsPayload {
	var tickets TicketsPayload
	data, _ := api.ListTickets(credentials, endpoint)
	//fmt.Println(data)

	err := json.Unmarshal([]byte(data), &tickets)

	if err != nil {
		fmt.Println(err)
	}

	//fmt.Printf("%+v\n", tickets)

	return tickets

}
