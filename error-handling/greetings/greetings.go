package greetings

import (
	"errors"
	"fmt"
)

// Returns a greeting for the given name
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message

	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome.", name)
	return message, nil // nil means No errors, in case of successful return, as the second value
}
