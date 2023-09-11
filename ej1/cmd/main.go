/*
Crear un programa que cumpla los siguiente puntos:
1. Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.
2. Tener un slice global de Product llamado Products instanciado con valores.
3. 2 métodos asociados a la estructura Product: Save(), GetAll(). El método Save() deberá tomar el slice de Products y añadir el producto desde el cual se llama al método. El método GetAll() deberá imprimir todos los productos guardados en el slice Products.
4. Una función getById() al cual se le deberá pasar un INT como parámetro y retorna el producto correspondiente al parámetro pasado.
5. Ejecutar al menos una vez cada método y función definido desde main().
*/

package main

import "fmt"

type Producto struct {
	Id          int
	Name        string
	Price       float64
	Description string
	Category    string
}

func (e Producto) save(p Producto) {
	products = append(products, p)
}

func (e Producto) getAll() {
	for _, product := range products {
		fmt.Println("Id:", product.Id)
		fmt.Println("Nombre:", product.Name)
	}
}

func (e Producto) getById(id int) (busquedaProducto Producto) {
	for _, product := range products {
		if product.Id == id {
			return product
		}
	}
	return
}

func (e Producto) printSingleProduct(prod Producto) {
	fmt.Println("Id:", prod.Id)
	fmt.Println("Nombre:", prod.Name)
}

var products = []Producto{
	{
		Id:          1,
		Name:        "iPhone 14 pro",
		Price:       150,
		Description: "-",
		Category:    "iPhones",
	},
	{
		Id:          2,
		Name:        "Samsung S22",
		Price:       140,
		Description: "-",
		Category:    "Samsungs",
	},
}

func main() {
	var temp Producto

	fmt.Println("Testing save() function:")
	temp.save(Producto{Id: 3,
		Name:        "Xiaomi nn",
		Price:       80,
		Description: "-",
		Category:    "Xiaomi"})

	fmt.Println("")
	fmt.Println("Testing getAll() function")
	temp.getAll()

	fmt.Println("")
	fmt.Println("Test getById() function")
	temp.printSingleProduct(temp.getById(2))
	temp.printSingleProduct(temp.getById(1))
}
