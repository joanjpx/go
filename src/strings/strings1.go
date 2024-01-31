package main

import "fmt"

func main() {
	s1 := "Hola mundo"
	s2 := `Esta es una linea 
	partida`

	fmt.Println(s1)
	fmt.Printf("El tipo de s1 es: %T\n",s1)
	fmt.Println(s2)
	fmt.Printf("El tipo de s2 es: %T\n",s2)

	fmt.Println("")

	bs := []byte(s1)

	fmt.Println(bs)

	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s1); i++ {
		
		// print the unicode 
		fmt.Printf("%#U\n", s1[i])
	}

	for i, v := range s1 {
		
		fmt.Printf("En el indice %d el valor es %v y el caracter es %c \n",i,v, v)
		// fmt.Printf("%U\n", r)
	}
}