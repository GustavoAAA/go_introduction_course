package main

import (
	"fmt"
	"os"
)

func main() {
	pl := fmt.Println
	file, err := os.Open("miFile.txt")
	if err != nil {
		pl(err)
		//El programa finaliza y devuelve un status que nosotros definimos
		os.Exit(1)
	}

	data := make([]byte, 100)
	c, err := file.Read(data)
	if err != nil {
		pl(err)
	}

	fmt.Printf("read %d bytes: %q\n", c, data[:c])

	env1 := os.Getenv("USERNAME")
	pl(env1)

	//Creo una variable de entorno
	os.Setenv("MI_ENV", "Valor de prueba")
	pl(os.Getenv("MI_ENV"))

	//Obtengo la ruta del directorio donde me encuentro
	dir, _ := os.Getwd()
	pl(dir)

	//Nos permite saber si una variable de entorno existe
	// si no existe la variable marca false en existe, sino muestra el valor
	env, existe := os.LookupEnv("ENV_TEST")
	pl(env, existe)

	env2, existe2 := os.LookupEnv("MI_ENV")
	pl(env2, existe2)

	//Simulamos los datos de conexión de una MongoDB como ejemplo
	//Las variables de entorno sirven para almacenar datos de configuración y de conexión
	os.Setenv("DB_USERNAME", "nahuel")
	os.Setenv("DB_PASSWORD", "sarasa")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "27018")
	os.Setenv("DB_NAME", "users")

	//Remplaza ${var} o $var por los valores que estan en las variables de entorno
	//Sirve para generar una cadena de conexión de una base de datos por ejemplo
	dbURL := os.ExpandEnv("mongodb://${DB_USERNAME}:${DB_PASSWORD}@$DB_HOST:$DB_PORT/$DB_NAME")
	fmt.Println(dbURL)
}
