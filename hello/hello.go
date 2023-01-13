package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetFlags(0)

	message, err := greetings.Hello("Cauã")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(message)

}
