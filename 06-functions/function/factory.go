package function

/*
Patron de diseño Factory + Strategy
Factory: Le pides una operación (SUM, SUB, DIV, MUL), la fábrica devuelve la función seleccionada
Strategy: Cada operación (sum, sub, div, mul) es una estrategia diferente para resolver un mismo problema
*/
func FactoryOperation(op Operation) func(float64, float64) float64 {

	switch op {
	case SUM:
		return sum
	case SUB:
		return sub
	case DIV:
		return div
	case MUL:
		return mul
	}

	return nil
}

func sub(a, b float64) float64 {
	return a - b
}

func sum(a, b float64) float64 {
	return a + b
}

func mul(a, b float64) float64 {
	return a * b
}

func div(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}
