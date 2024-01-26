package main

import "fmt"

// types and subtypes
type (
	A1 int
	A2 A1
)

type (
	B1 string
	B2 B1
	B3 []B1
	B4 B3
)

var x A1
var y A2

var z B1
var w B2
var v B3
var u B4

func main() {

	x = 42
	y = 43
	fmt.Println(x)
	fmt.Println(y)

	z = "James"
	w = "Bond"
	fmt.Println(z)
	fmt.Println(w)
	
	// set value for v as slices of B1
	v = []B1{z, B1(w)}
	fmt.Println(v)
	fmt.Printf("%T\n", v)
	
	// set value for u as a slice of B1
	u = []B1{z, B1(w)}
	fmt.Println(u)
	fmt.Printf("%T\n", u)

}