package main

import "fmt"

func main() {
	var number int
	var slice []int
	for {
		fmt.Scanln(&number)
		if number == 0 {
			break
		} else {
			slice = append(slice, number)
		}
	}
	fmt.Println(slice)
}
