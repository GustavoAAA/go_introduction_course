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

	x, y := function.Split(20)
	fmt.Println("Valor x:", x, "Valor y:", y)

	//Con esta función puedo poner en parámetros la cantidad de valores que quiera desde que sea número
	v = function.MSum(23, 12, 32, 12, 3, 1, 2, 3, 2, 1, 23, 12, 1)
	fmt.Println("Suma dinámica:", v)
	fmt.Println()

	v, err = function.MOperations(function.SUM, 2, 7, 1)
	fmt.Println("multy sum: ", v, " - error: ", err)

	v, err = function.MOperations(function.MUL, 2, 1, 3, 2, 1)
	fmt.Println("multy mul: ", v, " - error: ", err)

	v, err = function.MOperations(function.DIV, 2, 0, 1, 2, 1)
	fmt.Println("multy div: ", v, " - error: ", err)

	/*
		El comportamiento de esta función con un ejemplo es:
		Imagina que entras a una fábrica de herramientas y cuando llegas, te preguntan:
		¿Qué herramienta necesitas? ¿Martillo, destornillador, serrucho o llave?
		y según lo que respondes, la fábrica te entrega la herramienta correcta:
		* Pides martillo → te dan un martillo
		* Pides destornillador → te dan un destornillador
		Una vez que tienes la herramienta, ya puedes hacer lo que esa herramienta sabe hacer.
		* Con el martillo → golpear
		* Con el destornillador → atornillar
		* Con la llave → ajustar

		En resumen, la fábrica te da la herramienta y la herramienta hace la tarea

		En este caso lo primero que se hace es especificar que se quiere SUMAR,
		ya luego se le pasan los valores que se quieren sumar
	*/
	fn := function.FactoryOperation(function.SUB)
	v = fn(2, 3)
	fmt.Println("Sum:", v)

	fn = function.FactoryOperation(function.MUL)
	v = fn(2, 3)
	fmt.Println("Mul:", v)
}
