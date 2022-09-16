package main

import "strings"

func UserInputValidation(firstname string, lastname string, email string, usertickets uint, remainingtickets uint) (bool, bool, bool) {
	isValidName := len(firstname) > 2 && len(lastname) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := usertickets <= (remainingtickets) && usertickets > 0

	return isValidName, isValidEmail, isValidTickets
}
