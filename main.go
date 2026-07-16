package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remaingTickets uint = 50
	fmt.Printf("conference name %T, conferenceTickets %T, remaingTickets %T\n", conferenceName, conferenceTickets, remaingTickets)

	fmt.Printf("Welcome to %v booking application", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remaingTickets)
	fmt.Println("Get your ticket here to attend")

	// fmt.Println(conferenceName)

	fmt.Println(remaingTickets)
	// fmt.Println(&remaingTickets)

	//data take user
	var bookings []string //slices
	for {
		var firstName string
		var lastName string
		var email string
		var userTicket uint
		// var bookings [50]string//arrays

		fmt.Println("Enter your firstName :")
		fmt.Scan(&firstName)

		fmt.Println("Enter your lastName :")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email :")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets :")
		fmt.Scan(&userTicket)

		if remaingTickets > 0 {

			remaingTickets = remaingTickets - userTicket

			// bookings[52] = firstName + " " + lastName

			bookings = append(bookings, firstName+" "+lastName)
			fmt.Printf("The whole slice array: %v\n", bookings)
			// fmt.Printf("The first slice value: %v\n", bookings[0])
			fmt.Printf("Array slice type: %T\n", bookings)
			fmt.Printf("Array slice length: %v\n", len(bookings))

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				var firstName = names[0]
				firstNames = append(firstNames, firstName)
			}

			fmt.Printf("The first names of people Booking: %v\n", firstNames)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation at %v\n", firstName, lastName, userTicket, email)
			fmt.Printf("%v tickets remaining for %v\n", remaingTickets, conferenceName)

			boolenCheckingCase := remaingTickets == 0
			if boolenCheckingCase {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

		} else if remaingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
		} else {
			fmt.Printf("Our conference is booked out. Come back next year.")
		}

	}
}
