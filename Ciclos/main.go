package main

import "fmt"

func main() {

	fmt.Println("Bienvenidos a la clase de ciclos")

	fmt.Println("*******Bucle Normal**********")

	for i := 0; i < 11; i++ {
		fmt.Println(i)
	}
	fmt.Println("*******Bucle Infinito**********")

	for {
		println("Infinito")
		break
	}

	for rango := range [10]int{} {
		fmt.Println(rango)
	}

}
