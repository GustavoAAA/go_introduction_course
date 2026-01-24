package main

import (
	"fmt"

	"gitub.com/GustavoAAA/go_introduction_course/exercises/seccion3-35/matrix"
)

func main() {
	m, err := matrix.New([]float64{2, 7, 8}, []float64{4, 4, 7}, []float64{5, 6, 1})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	m.Print()
}
