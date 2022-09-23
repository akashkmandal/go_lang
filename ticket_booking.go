package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string
	conferenceName = "GoLang Conference"
	//Constant value assiged below, which is immutable
	const confereceTickets = 50

	var remainingTickets uint = 50
	var bookings []string
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", confereceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Grab your tickets from here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		//ask user for their name

		fmt.Println("Please enter below details to proceed with bookings...")
		fmt.Print("First Name: ")
		fmt.Scan(&firstName)
		fmt.Print("Last Name: ")
		fmt.Scan(&lastName)
		fmt.Print("Email Address: ")
		fmt.Scan(&email)
		fmt.Println("How many tickets do you need? ")
		fmt.Scan(&userTickets)

		if userTickets < remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)
			fmt.Println(bookings)
			fmt.Printf("Only %v tickets are left. Hurry Up and book your ticktes now.\n", remainingTickets)

			//creating slice
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Println("The first names of bookings are ", firstNames)

		} else if userTickets == remainingTickets {
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		} else {
			fmt.Printf("Sorry! We don't have %v tickets left. We only have %v tickets left.\n", userTickets, remainingTickets)
			continue
		}

	}
}
