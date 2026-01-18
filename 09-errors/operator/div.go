package operator

func Div(x, y float64) float64 {
	if y == 0 {
		//Estamos forzando un error, es como throw que pasa el error al main y lo toma en el defer
		panic("No se puede realizar división por cero")
	}
	return x / y
}
