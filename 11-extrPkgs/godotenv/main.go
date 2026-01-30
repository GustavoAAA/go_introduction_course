package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	p := fmt.Println

	p(os.Getenv("MY_ENV1"))
	p(os.Getenv("MY_ENV2"))
	p(os.Getenv("MY_ENV3"))

	//Toma el .env y registra los environment que se encuentren ahi
	//Existen solo mientras el programa esta corriendo, por lo que mueren cuando termina el programa
	//No afectan otros procesos
	if err := godotenv.Load(); err != nil {
		p(err)
	}

	//Se puede definir la ruta donde estan los environments que se deben crear
	if err := godotenv.Load("otherFolder/.var"); err != nil {
		p(err)
	}

	p(os.Getenv("MY_ENV1"))
	p(os.Getenv("MY_ENV2"))
	p(os.Getenv("MY_ENV3"))
	p(os.Getenv("MY_ENV4"))

	//Generar un map con los environment
	myEnv, err := godotenv.Read("otherFolder/.var")
	p(err)
	p(myEnv)

	//Vuelve a cargar los environment si se necesita mas adelante
	if err := godotenv.Overload(); err != nil {
		p(err)
	}
	p(os.Getenv("MY_ENV1"))
	p(os.Getenv("MY_ENV2"))
	p(os.Getenv("MY_ENV3"))
}
