package cli

import (
	"fmt"
	"github.com/linda-lai/golang-ticket-viewer/internal/zendesk"
)

var barrier = "|---------------------------------------------------------------------------------------------------------|"

func ShowTicket(payload zendesk.TicketPayload) string {

	data := fmt.Sprintf("%v\n| id:%v\n| subject:%v\n%v\ndescription:%v\n%v\n", barrier, payload.Ticket.ID, payload.Ticket.Subject, barrier, payload.Ticket.Description, barrier)

	return data

}

func ViewAllTickets(payload zendesk.TicketsPayload) string {
	values := map[string]string{"ticketID": "TICKET ID", "subject": "SUBJECT", "createdAt": "CREATED AT"}
	header := fmt.Sprintf("| %-10v || %-55v || %-31v|", values["ticketID"], values["subject"], values["createdAt"])

	var ticketDetails string

	for i := range payload.Tickets {
		ticketDetails += fmt.Sprintf("| %-10v || %-55v || %-30v |\n", payload.Tickets[i].ID, payload.Tickets[i].Subject, payload.Tickets[i].CreatedAt)
	}

	data := fmt.Sprintf("%v\n%v\n%v\n%v%v\n| TOTAL: %v ||", barrier, header, barrier, ticketDetails, barrier, payload.Count)

	return data
}