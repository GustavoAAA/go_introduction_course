package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	//Mensaje normal con la fecha y hora de ejecución
	log.Println("Log")

	//Imprime el mnsaje por pantalla y finaliza el programa
	//log.Fatal("mi error")

	//Hace la misma función del log.Fatal, pero muestra en donde sucede y en que linea sucede
	//log.Panic("mi error 2")

	//Crear un archivo de Log
	date := time.Now()
	file, err := os.Create(fmt.Sprintf("%d.log", date.UnixNano()))
	if err != nil {
		panic(err)
	}

	l := log.New(file, "success: ", log.LstdFlags|log.Lshortfile)
	l.Println("test 1")
	l.Println("test 2")
	l.Println("test 3")
	l.Println("test 4")
	l.SetPrefix("error: ")
	l.Println("test 3")
	l.Println("test 4")
}
