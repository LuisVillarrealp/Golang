package main

import (
	"fmt"
	contadorvocales "funcionalidades/ContadorVocales"
	conversor "funcionalidades/Conversor"
)

func main() {
	var usrOpcion string
	fmt.Println("Bienvenido a este programa decide que quieres realizar")
	for {
		fmt.Println("Quieres hacer un cambio de moneda o contar vocales. Elige 1 o 2")
		fmt.Println("Para terminar pon cualquier numero.")
		fmt.Printf("Ingresa tu opcion: ")
		fmt.Scan(&usrOpcion)
		if usrOpcion == "1" {
			var usrOpcionMoneda string
			var usd float64
			fmt.Println("Conversor de Dolares")
			fmt.Println("1. Euros")
			fmt.Println("2. Libras")
			fmt.Println("3. Won")
			fmt.Println("1. Bitcoin")
			fmt.Printf("Ingresa a que moneda quieres cambiar: ")
			fmt.Scan(&usrOpcionMoneda)
			fmt.Printf("Ingresa que cantidad de dolares quieres convertir: ")
			fmt.Scan(&usd)
			switch usrOpcionMoneda {
			case "1":
				Euros := conversor.Euros(usd)
				fmt.Println("La conversion de Dolar a Euro es: ", Euros)
			case "2":
				Libras := conversor.Libras(usd)
				fmt.Println("La conversion de Dolar a Libra es: ", Libras)
			case "3":
				Won := conversor.Won(usd)
				fmt.Println("La conversion de Dolar a Won es: ", Won)
			case "4":
				Bitcoin := conversor.Btc(usd)
				fmt.Println("La conversion de Dolar a Bitcoin es: ", Bitcoin)
			default:
				break
			}
		} else if usrOpcion == "2" {
			var frase string
			fmt.Println("Contador de Vocales")
			fmt.Println("Ingresa tu frase para contar las vocales en tu frase: ")
			fmt.Printf("Tu Frase aqui: ")
			fmt.Scan(&frase)
			a, e, i, o, u := contadorvocales.ContarVocales(frase)
			fmt.Println("La letra a se repite: ", a, " veces")
			fmt.Println("La letra e se repite: ", e, " veces")
			fmt.Println("La letra i se repite: ", i, " veces")
			fmt.Println("La letra o se repite: ", o, " veces")
			fmt.Println("La letra u se repite: ", u, " veces")
		} else {
			break
		}
	}
}
