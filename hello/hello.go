package main

import (
	"fmt"

	"aarneharju.com/go/modules/greetings"
)

func main() {
	// Get a greeting message and print it
	message := greetings.Hello("Petteri")
	fmt.Println(message)
}
