package main

import (
	"fmt"
)

type nuevoTipo2 string

const (
	INGRESO  nuevoTipo2 = "IN"
	SALIDA   nuevoTipo2 = "OUT"
	DESCANSO nuevoTipo2 = "REST"
	REUNION  nuevoTipo2 = "MEETING"
)

func proceso(proc nuevoTipo2) {
	switch proc {
	case INGRESO:
		fmt.Println("usted esta ingresando")
	case SALIDA:
		fmt.Println("usted esta saliendo")
	case DESCANSO:
		fmt.Println("usted esta entrando a su descanso")
	case REUNION:
		fmt.Println("usted esta entrando a una reunión")
	}
}

func main() {
	type nuevoTipo string

	var edad nuevoTipo = "15"
	fmt.Println(edad)

	proceso(INGRESO)
	proceso(REUNION)
	proceso(DESCANSO)
	proceso(SALIDA)

	type Carro struct {
		brand string
		year  int
	}

	type Tractor struct {
		car      Carro
		capacity int
	}

	miCarro := Carro{
		brand: "Chevrolet",
		year:  2015,
	}
	miTractor := Tractor{
		car:      miCarro,
		capacity: 5,
	}
	fmt.Println(miTractor)
}
