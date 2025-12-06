package libraries

func Print(nombre string) string {
	return "Hola " + nombre
}

type Carrito struct {
	ID      uint32
	Nombre  string
	Color   string
	Modelo  uint16
	Regalos []Regalos
}

type Regalos struct {
	ID          uint32
	Descripcion string
}

func AddRegalos(arr_aregalos []Regalos, ID int, Descripcion string) []Regalos {
	arr_aregalos = append(arr_aregalos, Regalos{
		ID:          uint32(ID),
		Descripcion: Descripcion,
	})
	return arr_aregalos
}

type Cuna struct {
	ID            uint32
	Nombre        string
	Vehiculos     []Carrito
	CantidadCarga uint16
}

func (c *Cuna) CantVehiculos() {
	c.CantidadCarga = uint16(len(c.Vehiculos))
}

func AddCarritos(ID int, Nombre string) Cuna {
	return Cuna{
		ID:     uint32(ID),
		Nombre: Nombre,
	}
}
