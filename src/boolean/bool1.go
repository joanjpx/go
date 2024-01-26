package main

import "fmt"

var c bool
var d bool
func main() {
	a := true
	b := false

	c = a && b

	fmt.Println(c)
	
	d = a || b
	fmt.Println(d)
	fmt.Println(2>1)
}