package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	greetUsers(conferenceName)

	fmt.Printf("We have a total of %v  tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		//ask user for their name
		fmt.Print("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Print("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Print("Enter your email address:")
		fmt.Scan(&email)

		fmt.Print("Enter the number of tickets:")
		fmt.Scan(&userTickets)

		isValiedName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketsNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValiedName && isValidEmail && isValidTicketsNumber {

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These are all our bookings:%v\n", firstNames)

			if remainingTickets == 0 {
				//endprogram
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}
		} else {
			if !isValiedName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketsNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}

	}
}

func greetUsers(confName string) {
	fmt.Printf("Welcome to %v booking application\n", confName)
}
