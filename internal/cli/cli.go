package cli

import (
	"github.com/linda-lai/golang-ticket-viewer/internal/zendesk"
)

func ShowTicket(credentials, endpoint, ticketID string) zendesk.TicketPayload {
	data := zendesk.UnmarshalZendeskTicket(credentials, endpoint, ticketID)

	return data
}

func ListAllTickets(credentials, endpoint string) zendesk.TicketsPayload {
	data := zendesk.UnmarshalZendeskTickets(credentials, endpoint)
	return data
}