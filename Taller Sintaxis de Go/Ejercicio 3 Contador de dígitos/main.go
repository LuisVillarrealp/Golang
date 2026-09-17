package main

import "fmt"

func main() {
	var n int
	fmt.Println("Ingresa un numero entero positivo")
	fmt.Scanf("%d", &n)

	num := n
	var digitos int = 0
	var suma int = 0

	for num > 0 {
		suma += num % 10
		num /= 10
		digitos++
	}

	fmt.Println("Cantidad de dígitos:", digitos)
	fmt.Println("Suma de dígitos:", suma)

	switch digitos {
	case 1:
		fmt.Println("Número de una cifra")
	case 2:
		fmt.Println("Número de dos cifras")
	case 3:
		fmt.Println("Número de tres cifras")
	default:
		fmt.Println("Número de varias cifras")
	}
}
