package main

import (
	"fmt"

	"github.com/linda-lai/golang-ticket-viewer/internal/api"
	"github.com/linda-lai/golang-ticket-viewer/internal/config"
	"github.com/linda-lai/golang-ticket-viewer/internal/zendesk"

	"github.com/linda-lai/golang-ticket-viewer/internal/cli"
	//"github.com/linda-lai/golang-ticket-viewer/internal/zendesk"
)

var text = "Enim qui elit minim cillum qui nisi deserunt ex excepteur nulla ex sint ut adipisicing. Lorem labore nulla non id aliquip ex excepteur est excepteur. Eu reprehenderit culpa consequat voluptate ullamco aute consequat.\n\nReprehenderit laborum deserunt minim exercitation anim officia ullamco duis anim. Officia adipisicing cillum aliquip exercitation do. Deserunt velit aute excepteur sit elit consequat reprehenderit occaecat nostrud quis consectetur ut. Voluptate mollit reprehenderit veniam qui cillum duis commodo exercitation enim cupidatat sunt voluptate velit non. Id pariatur aliqua in ipsum anim culpa non consectetur occaecat ut. Ex ex adipisicing ut sint mollit nisi consequat aute excepteur."

func main() {
	// Initialise config
	config := config.New()
	auth := config.BasicAuth
	endpoint := api.GetUrl(config.Subdomain)

	// Show Ticket By ID
	ticketID := "20"
	ticket := zendesk.UnmarshalZendeskTicket(
		auth,
		endpoint,
		ticketID)
	fmt.Println(cli.ShowTicket(ticket))

	// List All Tickets
	var tickets = zendesk.UnmarshalZendeskTickets(
		auth,
		endpoint)
	fmt.Println(cli.ViewAllTickets(tickets))

	fmt.Println(cli.WordWrap(text, 20))

	//fmt.Print("Enter text: \n")
	//var input string
	//_, err := fmt.Scanln(&input)
	//if err != nil {
	//	fmt.Fprintln(os.Stderr, err)
	//	return
	//}
	//fmt.Println(input)
}
