package operaciones

import "fmt"

func Suma(a, b int) int {
	return a + b
}

func SumaResta(y, z int) (int, int) {
	var suma int = y + z
	var resta int = 0
	if z > y {
		fmt.Println("Resta invalida")
	} else {
		resta = y - z
	}
	return suma, resta
}
func Sumatoria(numeros ...int) int {
	var suma int = 0
	for _, numero := range numeros {
		suma += numero
	}
	return suma
}
