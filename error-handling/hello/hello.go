package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Setting up a logger
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("")
	// if an error was returned, print it, the exit
	if err != nil {
		log.Fatal(err)
	}

	// if no error was returned, print the message
	fmt.Println(message)
}
