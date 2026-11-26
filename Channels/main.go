package main

import "fmt"

func calcularCuadrado(numero int, canal chan int) {
	res := numero * numero
	canal <- res
}
func main() {
	canal := make(chan int)
	go calcularCuadrado(4, canal)

	resultado := <-canal
	fmt.Println("El resultado es:", resultado)

}
