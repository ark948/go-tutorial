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

	// A slice of names
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting for names
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// if no error, print the map
	fmt.Println(messages)
}
