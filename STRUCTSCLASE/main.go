package main

import "fmt"

type Producto struct {
	Nombre string
	precio float64
	stock  int
}

// Funcion usando un Struct
func calcularValorInventario(p Producto) float64 {
	//multiplicar la cantidad por el precio
	return p.precio * float64(p.stock)
}
// METODOS CON PUNTERO
func (p *Producto) Vender(cantidad int) (bool, int) {
	if cantidad < 0 || cantidad > p.stock {
		fmt.Println("No existe suficiennte stock")
		return false, cantidad
	}
	fmt.Println("Si existe suficiennte stock")
	p.stock -= cantidad
	fmt.Println(p.stock)

	return true, cantidad
}
// Constructor 
func NewProduct(nombre string, precio float64, stock int) Producto{
	return Producto{
		Nombre: nombre,
		precio: precio,
		stock: 0,
	}

}
func (p *Producto) NuevoStock(cantidad int){
	if cantidad > 0{
		p.stock+=cantidad
	}
}

func main() {
	product := Producto{}

	product.Nombre = "Intel Core Ultra 7"
	product.precio = 850.90
	product.stock = 200

	product2 := Producto{
		Nombre: "Teclado",
		precio: 30.5,
		stock:  15,
	}
	fmt.Println("Producto 1: \n", product)
	fmt.Println("Producto 2: \n", product2)
	fmt.Println("El producto: ", product2.Nombre, "\nTiene un valor del inventario de: ", calcularValorInventario(product2))
	si, cant := product2.Vender(201)
	fmt.Println("Se vendio un producto: ", si, "La cantidad de :", cant)
	product3:=NewProduct("Mouse", 12.91, 50)
	fmt.Println("Producto 3: \n", product3)
	
}