package main

import (
	"fmt"
)

func main() {
	licence := true
	age := 10

	if age > 15 && licence {
		fmt.Println("Puede seguir avnzando")
	} else if age <= 15 || !licence {
		fmt.Println("No puede seguir circulando")
	}
}
