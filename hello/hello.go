package main

import (
	"fmt"

	"rsc.io/quote"

	"example.com/greetings"
)

func main() {
	quote.Go()
	fmt.Println(quote.Go())
	message := greetings.Hello("amit")

	fmt.Println(message)
}
