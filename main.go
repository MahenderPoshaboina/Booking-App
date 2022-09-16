package main

import (
	// "booking-app"
	"fmt"
	"strconv"
)

const conferenceTickets uint = 50

var conferencename = "Go conference"
var Remainingtickets uint = 50
var bookings = make([]map[string]string, 0)

func main() {

	greetUsers()

	for Remainingtickets > 0 && len(bookings) < 50 {

		firstname, lastname, email, usertickets := GetUserInput()
		isValidName, isValidEmail, isValidTickets := UserInputValidation(firstname, lastname, email, usertickets, Remainingtickets)
		if isValidName && isValidEmail && isValidTickets {

			BookTicket(firstname, lastname, usertickets, email)

			//Print first names
			firstNames := GetFirstNames()
			fmt.Printf("These are all our bookings %v \n", firstNames)

			fmt.Printf(" %v tickets are available\n", Remainingtickets)

			if Remainingtickets == 0 {
				fmt.Println(" Tickets Over.. Better Luck next time")
				break
			}
		} else {
			if !isValidName {
				fmt.Println(" First name or Last name may be short or Both may be short")
			}
			if !isValidEmail {
				fmt.Println(" Email doesn't contains @ sign")
			}
			if !isValidTickets {
				fmt.Println(" Number of tickets you entered is invalid ")
			}
		}

	}

}
func greetUsers() {
	fmt.Printf("Welcome to %v booking application \n", conferencename)
	fmt.Printf(" We have total of %v tickets and %v available\n", conferenceTickets, Remainingtickets)
	fmt.Println("Get your tickets here to attend")
}

func GetFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		firstNames = append(firstNames, booking["FirstName"])
	}
	return firstNames

}

func GetUserInput() (string, string, string, uint) {
	var firstname string
	var lastname string
	var email string
	var usertickets uint
	fmt.Printf("Enter Valid First Name: ")
	fmt.Scan(&firstname)
	fmt.Printf("Enter Valid Last Name: ")
	fmt.Scan(&lastname)
	fmt.Printf("Enter a Valid Email: ")
	fmt.Scan(&email)

	fmt.Printf("enter number of tickets: ")
	fmt.Scan(&usertickets)

	return firstname, lastname, email, usertickets
}

func BookTicket(firstname string, lastname string, usertickets uint, email string) {

	var UserData = make(map[string]string)
	UserData["FirstName"] = firstname
	UserData["LastName"] = lastname
	UserData["Email"] = email
	UserData[" Number of Tickets "] = strconv.FormatUint(uint64(usertickets), 10)

	bookings = append(bookings, UserData)
	fmt.Printf(" List of bookings is/are %v\n", bookings)
	Remainingtickets = Remainingtickets - uint(usertickets)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation mail at %v \n", firstname, lastname, usertickets, email)

}
