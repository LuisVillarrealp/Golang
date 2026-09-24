package main

import (
	"fmt"
	"practica/operaciones"
	"practica/saludo"
)

func main() {
	fmt.Println("👌👌 Bienvenid@s a la clase de Paquetes 😁😁")
	mensaje := saludo.Saludar("Luis")
	fmt.Println(mensaje)
	suma := operaciones.Suma(5, 6)
	suma1, resta := operaciones.SumaResta(6, 5)
	sumaVariadica := operaciones.Sumatoria(7, 8, 9, 4, 5, 6, 1)
	fmt.Println(suma)
	fmt.Println(suma1, resta)
	fmt.Println(sumaVariadica)
}
