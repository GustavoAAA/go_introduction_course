package main

import "fmt"

type Vehiculo interface {
	Mover()
}

type Auto struct{}

func (a Auto) Mover() {
	fmt.Println("El auto se mueve")
}

func HacerMover(v Vehiculo) {
	v.Mover()
}

func main() {
	auto := Auto{}
	HacerMover(auto)
}
