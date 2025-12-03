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

/*
Cuando se necesita que la función retorne 2 valores se debe colocar en paréntesis así:
(int, string)
*/
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

/*
Existe la posibilidad de poner las variables que se van a retornar,
en este caso el ejemplo es (x, y int) que se interpreta como:
Se retorna el valor que tenga "x" y "y" y ya no es necesario adicionarlos en el return
*/
func Split(v int) (x, y int) {
	x = v * 4 / 9
	y = v - x
	return
}

// Esta función recibe una cantidad dinámica de parámetros, se identifica con "..."
func MSum(values ...float64) float64 {
	var sum float64
	for _, v := range values {
		sum += v
	}
	return sum
}

func MOperations(op Operation, values ...float64) (float64, error) {

	if len(values) == 0 {
		return 0, errors.New("no hay valores para procesar")
	}
	sum := values[0]

	for _, v := range values[1:] {

		switch op {
		case SUM:
			sum += v
		case SUB:
			sum -= v
		case DIV:
			if v == 0 {
				return 0, errors.New("el valor no puede ser cero")
			}
			sum /= v
		case MUL:
			sum *= v
		}
	}
	return sum, nil
}
