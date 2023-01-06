package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v  tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int
	//ask user for their name
	fmt.Print("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email address:")
	fmt.Scan(&email)
	fmt.Print("Enter the number of tickets:")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v", firstName, lastName, userTickets, email)
}
