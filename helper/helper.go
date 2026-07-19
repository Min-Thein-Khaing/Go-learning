package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTicket uint , remaingTickets uint) (bool, bool, bool) {
	validateName := len(firstName) >= 2 && len(lastName) >= 2
	validateEmail := strings.Contains(email, "@")
	validateTicket := userTicket > 0 && userTicket <= remaingTickets
	return validateName, validateEmail, validateTicket
}