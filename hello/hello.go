package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello Niranjana")
	fmt.Println(quote.Go())
	var message string
	message, err := greetings.Hello("Gayatri")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
	message, err = greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
