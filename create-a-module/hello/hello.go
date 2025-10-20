package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("r4mn0")
	fmt.Println(message)
}

