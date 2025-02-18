package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ") //before every logs it will print
	// log.SetFlags(0) // if 0 thats mean no time and date

	message, err := greetings.Hello("Ashif")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
