package main

import "fmt"

/*

func <nombre>(param1, param2, ....param n)<valores de retorno>{
   ----------------------
   ----------------------
   ----------------------
   //return en el caso de que nuestra funcion retorne valores
}


*/

func saludar() {
	fmt.Println("Hola esta es mi primera funcion")
}
func bienvenida(nombre string) {
	fmt.Println("Bienvenid@", nombre)
}
func main() {
	var usr string
	fmt.Println("Ingresa Tu nombre")
	fmt.Scan(&usr)
	saludar()
	bienvenida(usr)
	suma, resta := sumaResta(4, 5)
	fmt.Println("suma=", suma, "resta=", resta)
	mostrarNum(5, 10, 15, 20, 25)
	fmt.Println("La sumatoria es: ", sumarNum(5, 10, 15, 20, 25))
}

func sumaResta(y, z int) (int, int) {
	var suma int = y + z
	var resta int = 0
	if z > y {
		fmt.Println("Resta invalida")
	} else {
		resta = y - z
	}
	return suma, resta
}

/* Variadic Functions ======== Funciones Variadicas */
func mostrarNum(numeros ...int) {
	fmt.Println("Los numeros ingresados son: ", numeros)
}
func sumarNum(numeros ...int) int {
	var suma int = 0
	for _, numero := range numeros {
		suma += numero
	}
	return suma
}
