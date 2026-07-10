package main

import "fmt"

func main() {
	var conferenceName  = "Go Conference"
	const conferenceTickets  = 50
	var remaingTickets  = 50
	fmt.Printf("conference name %T, conferenceTickets %T, remaingTickets %T\n", conferenceName, conferenceTickets, remaingTickets)


	fmt.Printf("Welcome to %v booking application", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remaingTickets)
	fmt.Println("Get your ticket here to attend")

	// fmt.Println(conferenceName)

	var userName string
	var userTicket int

	fmt.Printf("User %v book %v tickets", userName, userTicket)

}
