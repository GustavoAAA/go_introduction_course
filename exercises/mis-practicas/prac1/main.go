package main

import (
	"fmt"
	"strings"
)

func esPalindromo(texto string) bool {
	prueba2 := strings.Split(strings.ToLower(strings.ReplaceAll(texto, " ", "")), "")
	var slice1 []string

	for i := len(prueba2) - 1; i >= 0; i-- {
		slice1 = append(slice1, prueba2[i])
	}

	if len(prueba2) != len(slice1) {
		return false
	}
	fmt.Println(prueba2)
	fmt.Println(slice1)

	for v := range prueba2 {
		if prueba2[v] != slice1[v] {
			return false
		}
	}
	return true
}

func main() {

	res := esPalindromo("Anita lava la tina")
	if res {
		fmt.Println("Es palíndromo")
	} else {
		fmt.Println("No es palíndromo")
	}

}

/*
func esPalindromo(s string) bool {
	// Convertir a minúsculas
	s = strings.ToLower(s)

	// Filtrar solo letras y números
	var cleaned []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			cleaned = append(cleaned, r)
		}
	}

	// Verificar palíndromo
	i, j := 0, len(cleaned)-1
	for i < j {
		if cleaned[i] != cleaned[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func main() {
	texto := "Anita lava la tina"
	if esPalindromo(texto) {
		fmt.Println("Es palíndromo")
	} else {
		fmt.Println("No es palíndromo")
	}
}
*/
