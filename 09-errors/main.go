package main

import (
	"errors"
	"fmt"

	"gitub.com/GustavoAAA/go_introduction_course/09-errors/operator"
)

func main() {
	var err error
	err = errors.New("Mi primer error")
	// Se pone .Error() cuando se necesita pasar la variable de tipo error a string
	fmt.Println(err.Error())

	err2 := fmt.Errorf("mi error formateado, string: %s, number: %d", "mi texto", 23)
	fmt.Println(err2)

	//Es una función anónima, hace una tarea puntual dentro de otra función (también llamada Immediately Invoked Function Expression, IIFE).
	func() {
		fmt.Println("función que se ejecuta de inmediato en el punto que se tenga")
	}()

	// Cuando se registra como defer se ejecutará cuando la función que la contiene termine
	//Ejemplos de uso de esta función:
	//Limpieza de recursos como cerrar un archivo (file, _ := os.Open("test.txt") defer file.Close())
	//Logs al finalizar la ejecución de una func
	//Medir tiempos de ejecución (start := time.Now() defer func() { fmt.Println(time.Since(start)) }())
	defer func() {
		//Ejecuta al final
		//recover() captura el error o excepción que salga en la ejecución
		if r := recover(); r != nil {
			fmt.Println("error recuperado: ", r)
		}
	}()

	// El panic sirve para lanzar un error, muy parecido a throw new Exception("algo salió mal");
	//panic("algo salió mal")

	z := operator.Div(4, 0)

	fmt.Println("Z es ", z)
}
