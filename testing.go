/*package main

import "fmt"

func main() {
	var pname string
	var lname string
	var email string

	var totalTickets uint = 50
	var bookedTickets uint
	var remainingTickets uint

	bookings := []string{}

	fmt.Print("What's your first name?\n")
	fmt.Scan(&pname)

	fmt.Print("What's your last name?\n")
	fmt.Scan(&lname)

	fmt.Print("What's your email address?\n")
	fmt.Scan(&email)

	fmt.Print("How many tickets would you like to buy?\n")
	fmt.Scan(&bookedTickets)

	remainingTickets = totalTickets - bookedTickets
	bookings = append(bookings, pname+" "+lname)
	for i := 0; i < len(bookings); i++ {
		fmt.Printf("%v booked %v tickets.\n", bookings[i], bookedTickets)
	}

	fmt.Printf("%v %v has booked %v tickets. Your email is: %v\n", pname, lname, bookedTickets, email)
	fmt.Printf("%v tickets are remaining.", remainingTickets)
}*/
