package function

import (
	"errors"
	"fmt"
)

type Operation int

const (
	SUM Operation = iota
	SUB
	DIV
	MUL
)

/*
Recordar que en la letra inicial de la función es:
Mayúscula: Público
Minúscula: Privado
*/
func Add(x int, y int) int {
	return x + y
}

func RepeatString(increment int, value string) {
	for i := 0; i < increment; i++ {
		fmt.Print(value)
	}
}

func Calc(op Operation, x, y float64) (float64, error) {
	switch op {
	case SUM:
		return x + y, nil
	case SUB:
		return x - y, nil
	case DIV:
		if y == 0 {
			return 0, errors.New("y mustn´t be zero")
		}
		return x / y, nil
	case MUL:
		return x * y, nil
	}

	return 0, errors.New("has been an error")
}
