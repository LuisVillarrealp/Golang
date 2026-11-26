package main

import (
	"fmt"
	"time"
)

func saludar(nombre string) {
	for i := 0; i < 4; i++ {
		fmt.Println("Hola: ", nombre)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go saludar("Jony")
	saludar("Luis")
}
