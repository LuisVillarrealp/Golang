package main

import "fmt"

func main() {
	numero := 10
	puntero := &numero
	fmt.Println("Valor: ", numero)
	fmt.Println("Direccion:", &numero)
	fmt.Println("Puntero: ", puntero)
	fmt.Println("Valor Puntero: ", *puntero)

}
