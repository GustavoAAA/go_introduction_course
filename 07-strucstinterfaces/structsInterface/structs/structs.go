package structs

/*
Los struct tags `json:"id"` son metadatos en forma de string que otras librerías
pueden leer mediante reflection.
*/
type Product struct {
	ID    uint     `json:"id"`
	Name  string   `json:"name"`
	Type  Type     `json:"type"`
	Count uint16   `json:"count"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

type Type struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

/*
(p Product) significa que TotalPrice() es un método asociado al tipo Product
*/
func (p Product) TotalPrice() float64 {
	return float64(p.Count) * p.Price
}

/*
Cuando lleva el * significa que es un puntero, tema de mas adelante
Pero si no se pone como puntero no cambia el valor
*/
func (p *Product) SetName(name string) {
	p.Name = name
}

func (p *Product) AddTags(tags ...string) {
	p.Tags = append(p.Tags, tags...)
}

type Car struct {
	ID       uint      `json:"id"`
	UserID   uint      `json:"user_id"`
	Products []Product `json:"products"`
}

func (c *Car) AddProducts(products ...Product) {
	c.Products = append(c.Products, products...)
}

func (c Car) Total() float64 {
	var total float64
	for _, c := range c.Products {
		total += c.TotalPrice()
	}
	return total
}

/*
Crea un carrito de compras, en este caso solo le asigna el UserID por parámetro y el ID como dato fijo
Para los productos nuevos del struct Car sería ejecutar la función AddProducts() que es un append de productos
*/
func NewCar(userID uint) Car {
	return Car{UserID: userID, ID: 1}
}
