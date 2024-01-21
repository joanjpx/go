package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
    // Get a greeting message and print it.
    // message := greetings.Hello("Gladys")
    
    fmt.Println(greetings.Hi("Vivian"))
    fmt.Println(greetings.Hello("Viv"))
    fmt.Println(greetings.Regards("Vivi"))
    fmt.Println(greetings.Bye("Vivian"))
}