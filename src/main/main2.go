package main

import "fmt"

var a int
var b string
var c bool

func main() {
	
	a := 42
	b := "James Bond"
	c := true

	fmt.Println(a, b, c)
	fmt.Printf("%T, %T, %T", a, b, c)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}