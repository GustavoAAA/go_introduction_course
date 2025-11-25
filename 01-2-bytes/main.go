package main

import "fmt"

func main() {
	var A byte = 'A'
	var a byte = '1'
	var R byte = 82
	var s byte = 115
	vector := []byte{65, 97, 82, 115}

	fmt.Println("Valor de la variable", A)
	fmt.Println("Valor de la variable", a)
	fmt.Println("Valor de la variable", string(R))
	fmt.Println("Valor de la variable", string(s))
	fmt.Println("Valor de la variable", string(vector))
}
