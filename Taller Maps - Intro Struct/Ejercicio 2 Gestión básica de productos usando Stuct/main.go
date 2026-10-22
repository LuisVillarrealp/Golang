package main

import "fmt"

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func main() {
	var producto1 Producto = Producto{Nombre: "Grand Theft Auto VI Ultimate Edition", Precio: 99.99, Cantidad: 0}
	var producto2 Producto = Producto{Nombre: "EA SPORTS FC 27 Ultimate Plus Edition", Precio: 149.99, Cantidad: 11}
	var producto3 Producto = Producto{Nombre: "God Of War Ragnarok", Precio: 69.99, Cantidad: 30}

	var productostotales []Producto = []Producto{producto1, producto2, producto3}

	fmt.Println("Bienvenido a la Gestion de Productos")
	fmt.Println("Tus productos son estos")
	for contador, videojuego := range productostotales {
		fmt.Println("	Producto numero ",contador+1 )
		mostrarProducto(videojuego)
	}
}

func mostrarProducto(videojuego Producto) {
	fmt.Println("Nombre: ", videojuego.Nombre)
	fmt.Println("Precio: ", videojuego.Precio)
	fmt.Println("Cantidad: ", videojuego.Cantidad)
	fmt.Println("Valor total: ", valorTotal(videojuego))
	if tieneStock(videojuego) {
		fmt.Println("Estado de stock: disponible")
	} else {
		fmt.Println("Estado de stock: sin stock")
	}
}

func valorTotal(videojuego Producto) float64 {
	return videojuego.Precio * float64(videojuego.Cantidad)
}

func tieneStock(videojuego Producto) bool {
	return videojuego.Cantidad > 0
}
