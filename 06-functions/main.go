package main

import (
	"fmt"

	"gitub.com/GustavoAAA/go_introduction_course/06-functions/function"
)

func main() {
	fmt.Println(function.Add(3, 4))

	function.RepeatString(10, "as")
	fmt.Println()

	v, err := function.Calc(function.SUM, 3, 6)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Value:", v)
	}

	v, err = function.Calc(function.DIV, 3, 0)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Value:", v)
	}
}
