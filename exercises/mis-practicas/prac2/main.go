package main

import (
	"fmt"

	"gitub.com/GustavoAAA/go_introduction_course/exercises/mis-practicas/prac2/libraries"
)

func main() {
	nombre := "Pedro"
	fmt.Println(libraries.Print(nombre))

	var arr_regalos []libraries.Regalos

	regalos := make(map[int]string)
	regalos[1] = "Kit de carreteras"
	regalos[2] = "Polarizado"
	regalos[3] = "Repuesto de llanta"

	for k, v := range regalos {
		arr_regalos = libraries.AddRegalos(arr_regalos, k, v)
	}

	chevrolet_aveo := libraries.Carrito{
		ID:      1,
		Nombre:  "Chevrolet",
		Color:   "Aveo",
		Modelo:  2010,
		Regalos: arr_regalos,
	}

	nissan_march := libraries.Carrito{
		ID:      2,
		Nombre:  "Nissan",
		Color:   "March",
		Modelo:  2015,
		Regalos: arr_regalos,
	}

	cuna := libraries.AddCarritos(1, "Cuna Cali")
	cuna.Vehiculos = append(cuna.Vehiculos, chevrolet_aveo)
	cuna.Vehiculos = append(cuna.Vehiculos, nissan_march)
	cuna.CantVehiculos()

	fmt.Println(cuna)

}
