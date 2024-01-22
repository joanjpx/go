package main

import (
	"fmt"

	"example.com/greetings"
)


var z int = 21

func main() {
    // Get a greeting message and print it.
    // message := greetings.Hello("Gladys")
    
    var message string = "Message";

    // call function from another package/module
    fmt.Println(greetings.Hi("Vivian"))
    fmt.Println(greetings.Hello("Viv"))
    fmt.Println(greetings.Regards("Vivi"))
    fmt.Println(greetings.Bye("Vivian"))
    fmt.Println(message)


    

    for i := 0; i < 10; i++ {
        
        fmt.Println(i);
    }

    Numero()
}


func Numero() {

    fmt.Println(z)
}