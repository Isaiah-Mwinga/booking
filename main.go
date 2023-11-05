package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 100

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We currently have %v tickets available!\n", conferenceTickets)
}