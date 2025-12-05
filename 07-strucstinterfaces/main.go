package main

import (
	"encoding/json"
	"fmt"

	"gitub.com/GustavoAAA/go_introduction_course/07-strucstinterfaces/structsInterface/structs"
)

func main() {

	//No es necesario poner todos los valores ya que no son obligatorios
	var p1 structs.Product
	p1.ID = 12
	p1.Name = "Detergente"
	p1.Price = 1200.0
	fmt.Println(p1)

	p2 := structs.Product{}
	p2.ID = 14
	p2.Name = "Aceite para freir"
	p2.Price = 2500
	fmt.Println(p2)

	//Forma de ingresar los valores, pero no recomendada
	p3 := structs.Product{2, "Pimienta", structs.Type{}, 0, 1500, nil}
	fmt.Println(p3)

	//Si se van a ingresar de inmediato la mejor manera es asi:
	p4 := structs.Product{
		ID:    3,
		Name:  "Sartén",
		Type:  structs.Type{},
		Price: 3200,
	}
	fmt.Println(p4)

	p5 := structs.Product{
		ID:   4,
		Name: "Heladera marca sarasa",
		Type: structs.Type{
			ID:          1,
			Code:        "A",
			Description: "Electrodoméstico",
		},
		Price: 3200,
		Tags:  []string{"heladera", "sarasa", "freezer", "refrigerador"},
		Count: 5,
	}
	fmt.Println(p5)

	/*
		Devuelve un array con todos los valores en código ASCII, el resultado de traducirlo da:
		{"id":4,"name":"Heladera marca sarasa","type":{"id":1,"code":"A","description":"Electrodoméstico"},"count":5,"price":3200,"tags":["heladera","sarasa","freezer","refrigerador"]}
	*/
	v, err := json.Marshal(p5)
	fmt.Println(err)
	fmt.Println(v)
	fmt.Println(string(v))

	fmt.Println("Precio total:", p5.TotalPrice())

	p5.SetName("Cambio de nombre")
	p5.AddTags("tag1", "tag2", "tag3")
	fmt.Println(p5)

	p7 := structs.Product{
		Name:  "Naranja",
		Price: 20,
		Type: structs.Type{
			Code:        "B",
			Description: "Alimento",
		},
		Tags:  []string{"alimento", "verdura"},
		Count: 20,
	}

	p8 := structs.Product{
		Name:  "Cortina",
		Price: 2700,
		Type: structs.Type{
			Code:        "C",
			Description: "Hogar",
		},
		Tags:  []string{"hogar", "cortina"},
		Count: 3,
	}

	car := structs.NewCar(11212)
	car.AddProducts(p5, p7, p8)

	fmt.Println("PRODUCTS CAR")
	fmt.Println("Total Products: ", len(car.Products))
	fmt.Printf("Total Car: $%.2f\n", car.Total())
	fmt.Println()
	fmt.Println(car)
}
