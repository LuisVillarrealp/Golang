package main

import "fmt"

func main() {
	type Producto struct {
		Nombre string
		precio float64
		stock  int
	}
	product := Producto{}

	product.Nombre = "Intel Core Ultra 7"
	product.precio = 850.90
	product.stock = 200

	product2 := Producto{
		Nombre: "Teclado",
		precio: 30.5,
		stock:  15,
	}
	fmt.Println(product)
	fmt.Println(product2)
}
func calcularStock(p Producto)float64{
//multiplicar la cantidad por el precio 

}
