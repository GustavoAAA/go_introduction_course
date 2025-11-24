package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	var myIntVar int
	myIntVar = -11
	fmt.Println("My variable is:", myIntVar)

	// variable uint, solo numeros positivos
	var myUintVar uint
	myUintVar = 100
	fmt.Println("My variable is:", myUintVar)

	// variable string
	var myStringVar string
	myStringVar = "Cadena de texto"
	fmt.Println("My variable is:", myStringVar)

	// variable booleana
	var myBoolVar bool
	myBoolVar = true
	fmt.Println("My variable is:", myBoolVar)

	// vemos la direccion de memoria de nuestra variable
	fmt.Println("La dirección en memoria de mi variable es:", &myStringVar)

	// go tomara el valor que asignamos e infiere tipo de dato (al usar :=)
	myIntVar2 := 12
	fmt.Println("My variable is", myIntVar2)

	myStringVar2 := "Cadena nueva de texto"
	fmt.Println("My variable is:", myStringVar2)

	// Se infiere el tipo de dato de la constante
	const myFirstConst = "a12"
	fmt.Println("My variable is:", myFirstConst)

	// especificamos el tipo de dato de la constante
	const myIntConst int = 12
	fmt.Println("My variable is:", myIntConst)

	/*
		int8	8 bits				-128 to 127
		int16	16 bits				-2^15 to 2^15 -1
		int32	32 bits				-2^31 to 2^31 -1
		int64	64 bits				-2^63 to 2^63 -1
		int		Platform dependent	Platform dependent
	*/

	fmt.Println()
	var my8bitsIntVar int8
	fmt.Printf("Int default value: %d\n", my8bitsIntVar)

	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my8bitsIntVar, unsafe.Sizeof(my8bitsIntVar), unsafe.Sizeof(my8bitsIntVar)*8)

	var my16bitsIntVar int16
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my16bitsIntVar, unsafe.Sizeof(my16bitsIntVar), unsafe.Sizeof(my16bitsIntVar)*8)

	var my32bitsIntVar int32
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my32bitsIntVar, unsafe.Sizeof(my32bitsIntVar), unsafe.Sizeof(my32bitsIntVar)*8)

	var my64bitsIntVar int64
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my64bitsIntVar, unsafe.Sizeof(my64bitsIntVar), unsafe.Sizeof(my64bitsIntVar)*8)

	/*
		Type	Size				Range
		uint8	8 bits				0 to 255
		uint16	16 bits				0 to 2^16 -1
		uint32	32 bits				0 to 2^32 -1
		uint64	64 bits				0 to 2^64 -1
		uint	Platform dependent	Platform dependent
	*/

	fmt.Println()
	var my8BitsUnitVar uint8
	fmt.Printf("Uint default value: %d\n", my8BitsUnitVar)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my8BitsUnitVar, unsafe.Sizeof(my8BitsUnitVar), unsafe.Sizeof(my8BitsUnitVar)*8)

	var my16BitsUnitVar uint16
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my16BitsUnitVar, unsafe.Sizeof(my16BitsUnitVar), unsafe.Sizeof(my16BitsUnitVar)*8)

	var my32BitsUnitVar uint32
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my32BitsUnitVar, unsafe.Sizeof(my32BitsUnitVar), unsafe.Sizeof(my32BitsUnitVar)*8)

	var my64BitsUnitVar uint64
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my64BitsUnitVar, unsafe.Sizeof(my64BitsUnitVar), unsafe.Sizeof(my64BitsUnitVar)*8)

	fmt.Println()

	var myFloast32Var float32
	fmt.Printf("Float default value: %f\n", myFloast32Var)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myFloast32Var, unsafe.Sizeof(myFloast32Var), unsafe.Sizeof(myFloast32Var)*8)
	fmt.Println()

	var myFloat64Var float64
	fmt.Printf("Float default value: %f\n", myFloat64Var)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myFloat64Var, unsafe.Sizeof(myFloat64Var), unsafe.Sizeof(myFloat64Var)*8)
	fmt.Println()

	// string es una secuencia de bytes
	var myStringVar3 string
	fmt.Printf("String default value: %s\n", myStringVar3)

	myStringVar5 := `My string variable in golang
	with multiple
	line`

	fmt.Printf("The variable value is: %s\n", myStringVar5)

	//Esto es un scope, osea que lo que yo defina aqui adentro solo sirve aqui adentro, no sirve en otra parte del código
	{
		// conversiones
		//Convertir de decimal a texto
		fmt.Println()
		floatVar := 33.11
		fmt.Printf("type: %T, value: %f\n", floatVar, floatVar)
		floatStrVar := fmt.Sprintf("%.2f", floatVar)
		fmt.Printf("type: %T, value: %s\n", floatStrVar, floatStrVar)

		//Convertir de int a texto
		intVar := 22
		fmt.Printf("type: %T, value: %d\n", intVar, intVar)
		intStrVar := fmt.Sprintf("%d", intVar)
		fmt.Printf("type: %T, value: %s\n", intStrVar, intStrVar)

		//Convertir de texto a int64
		intVal1, err := strconv.ParseInt("1333", 0, 64)
		fmt.Println(err)
		fmt.Printf("type: %T, value: %d\n", intVal1, intVal1)

		//Convertir de texto a int64, en este caso simulando un error
		intVal2, err := strconv.ParseInt("aa122", 0, 64)
		fmt.Println(err)
		fmt.Printf("type: %T, value: %d\n", intVal2, intVal2)

		//Convertir texto a decimal, el 64 es el tamaño
		//Se omite el error poniendo el "_"
		floatVar1, _ := strconv.ParseFloat("-11.2", 64)
		fmt.Printf("type: %T, value: %f\n", floatVar1, floatVar1)

		//Convertir de decimal a texto
		myFloatToText := strconv.FormatFloat(3.1416, 'f', -1, 32)
		// 'f': punto fijo - 'e': notación científica - 'g': elige entre 'f' y 'e' según el tamaño del número
		fmt.Println("Convertir de decimal a texto", myFloatToText)

	}

	fmt.Println()
}
