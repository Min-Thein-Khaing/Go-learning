package main

import (
	"booking-app/helper"
	"fmt"
	// "strconv"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets int = 50

var remaingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}
func main() {

	greetUser()
	// fmt.Println(conferenceName)

	fmt.Println(remaingTickets)
	// fmt.Println(&remaingTickets)

	//slices
	//data take user

	firstName, lastName, email, userTicket := getUserInput()
	validateName, validateEmail, validateTicket := helper.ValidateUserInput(firstName, lastName, email, userTicket, remaingTickets)

	if validateName && validateEmail && validateTicket {
		calcTicket(firstName, lastName, userTicket, email)
		wg.Add(1)
		go sendTicket(userTicket, firstName, lastName, email)

		
		
		//firstName
		firstNames := printFirstName()
		fmt.Printf("The first names of people Booking: %v\n", firstNames)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation at %v\n", firstName, lastName, userTicket, email)
		fmt.Printf("%v tickets remaining for %v\n", remaingTickets, conferenceName)

		boolenCheckingCase := remaingTickets == 0
		if boolenCheckingCase {
			fmt.Println("Our conference is booked out. Come back next year.")
			// break
		}

	} else if remaingTickets == 0 {
		fmt.Println("Our conference is booked out. Come back next year.")
	} else {
		fmt.Printf("Your validation is invalid. Please try again\n")
	}

	wg.Wait()

}

// city := "London"

// switch city {
// case "New York":
// //some code
// case "Singapore", "Hong Kong":
// //some code
// case "London", "Berlin":
// //some code

// //some code
// case "Mexico City":
// //some code

// //
// default:
// 	fmt.Println("Default case")

// }

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remaingTickets)
	fmt.Println("Get your ticket here to attend")
}

func printFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint
	// var bookings [50]string//arrays

	//validation

	fmt.Println("Enter your firstName :")
	fmt.Scan(&firstName)

	fmt.Println("Enter your lastName :")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email :")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets :")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

func calcTicket(firstName string, lastName string, userTicket uint, email string) (string, string, uint) {

	remaingTickets = remaingTickets - userTicket

	// bookings[52] = firstName + " " + lastName

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTicket,
	}

	bookings = append(bookings, userData)
	fmt.Printf("The whole slice array: %v\n", bookings)
	// fmt.Printf("The first slice value: %v\n", bookings[0])
	fmt.Printf("Array slice type: %T\n", bookings)
	fmt.Printf("Array slice length: %v\n", len(bookings))

	return firstName, lastName, userTicket
}

func sendTicket(userTicket uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets ,%v %v ", userTicket, firstName, lastName)
	fmt.Println("#######################")
	fmt.Printf("sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#######################")
	wg.Done()
}
