package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Bye(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("👍🏼, %v. goodbye!", name)
    return message
}