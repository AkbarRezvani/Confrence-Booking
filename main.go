package main

import (
	"fmt"

	"Booking-App/validation"
)

var confrenceName = "Go Confrence"

const confrenceTickets int = 50

var remainingTickets uint = 50

// slice
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	for {

		// asking for user's info
		firstName, lastName, email, userTickets := getUserInput()

		// validate user input

		isValidName, isValidEmail, isValidTicketNumber := validation.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			//book tickets in system

			bookTicket(userTickets, firstName, lastName, email)
			//print only first names

			fmt.Printf("The first names are: %v \n", getFirstNames())

			if remainingTickets == 0 {
				//end program
				fmt.Println("our confrence is booked out")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("firt name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}

			continue
		}

	}

}

func greetUsers() {
	fmt.Println("welcome to", confrenceName, "booking application")
	fmt.Printf("We have a total of %v and we have %v tickets availablee\n", confrenceTickets, remainingTickets)
	fmt.Println("click here to get your tickets")
}

func getFirstNames() []string {
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
	var userTickets uint
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email:")
	fmt.Scan(&email)
	fmt.Println("Enter the number of tickets you want:")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("the information of bookings are: %v \n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. with confirmation email: %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("remaining tickets are: %v \n", remainingTickets)

}
