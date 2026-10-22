package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var usrInputPalabra string
	var usrOpcion string
	var lector *bufio.Scanner = bufio.NewScanner(os.Stdin)

	fmt.Println("Bienvenido al analizador de texto")
	fmt.Println("Ingrese un texto o frase y el programa va analizar palabra por palabra")

	for {
		fmt.Println("Que desea hacer")
		fmt.Println("1. Ingresar una frase")
		fmt.Println("2. Terminar")
		fmt.Printf("Ingresa tu opcion: ")
		lector.Scan()
		usrOpcion = lector.Text()

		if usrOpcion == "1" {
			var palabras map[string]int = map[string]int{}

			fmt.Printf("Ingresa la palabra: ")
			lector.Scan()
			usrInputPalabra = lector.Text()

			var listaPalabras []string = strings.Fields(strings.ToLower(usrInputPalabra))

			for _, palabra := range listaPalabras {
				palabras[palabra]++
			}

			if len(palabras) == 0 {
				fmt.Println("Frase invalida no hay palabras")
			} else {
				fmt.Println("			RESULTADOS")
				for pal, repeticiones := range palabras {
					fmt.Println("La palabra ", pal, " se repite ", repeticiones, " veces")
				}
				masRepetida, veces := palabraMasRepetida(palabras)
				fmt.Println("La palabra que mas se repite es: ", masRepetida, " con ", veces, " veces")
			}

		} else if usrOpcion == "2" {
			break
		} else {
			fmt.Println("Opcion invalida, intenta de nuevo")
		}
	}
}

func palabraMasRepetida(palabras map[string]int) (string, int) {
	var masRepetida string
	var mayorRepeticiones int = 0
	for pal, repeticiones := range palabras {
		if repeticiones > mayorRepeticiones {
			masRepetida = pal
			mayorRepeticiones = repeticiones
		}
	}
	return masRepetida, mayorRepeticiones
}
