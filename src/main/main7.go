package main

import "fmt"

type numero int

var x numero
var y int
func main() {

	fmt.Println(x)
	fmt.Printf("El tipo de x es : %T\n", x)
	
	x = 42
	
	fmt.Println(x)
	fmt.Printf("El tipo de x es : %T\n", x)
	
	// convertion of x from numero type to int type
	y = int(x)

	fmt.Println(y)
}