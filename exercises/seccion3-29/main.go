package main

import "fmt"

func main() {
	var option string
	var slice []string
	code := map[string]string{
		"10": "notebook",
		"15": "tv",
		"21": "heladera",
		"27": "monitor",
		"35": "camara",
	}

	for {
		fmt.Scanln(&option)
		if option == "0" {
			break
		} else if _, existe := code[option]; existe == false {
			slice = append(slice, "No encontrado")
		} else {
			slice = append(slice, code[option])
		}
	}
	fmt.Println(slice)
}
