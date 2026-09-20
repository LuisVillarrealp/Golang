package main

import (
	"fmt"
)

func main() {
	var usrOpcion string
	fmt.Println("Bienvenidos a este programa puedes realizar las siguientes opciones.")
	fmt.Println("Cada que termines una opcion se te desplegara el menu con las opciones \n ")
	for {
		fmt.Println("			Menu")
		fmt.Println("La opcion 1 es para calcular la nota de estudiantes de un curso del 1 al 100")
		fmt.Println("La opcion 2 es para calcular la suma desde 1 al numero elegido")
		fmt.Println("La opcion 3 para convertir de Celsius a Farenheit")
		fmt.Println("La opcion 4 es para convertir de Farenheit a Celsius")
		fmt.Println("La opcion 0 o 'salir' termina el programa. \n ")
		fmt.Print("Ingresa tu opcion aqui: ")
		fmt.Scan(&usrOpcion)
		if usrOpcion == "1" {
			fmt.Print("Ingrese el numero de estudiantes: ")
			var numeroEstudiantes int
			fmt.Scan(&numeroEstudiantes)
			promedioGeneral := opcion1(numeroEstudiantes)
			fmt.Println("El promedio de notas es: ", promedioGeneral)
		} else if usrOpcion == "2" {
			fmt.Print("Ingresa un numero para calcular la suma: ")
			var n int
			fmt.Scan(&n)
			sumaNumFinal := opcion2(n)
			fmt.Println("La suma de los numeros: ", sumaNumFinal)
		} else if usrOpcion == "3" {
			fmt.Print("Ingresa una temperatura en Celsius para convertir a Farenheit: ")
			var temperaturaCelsius float64
			fmt.Scan(&temperaturaCelsius)
			Farenheit := opcion3(temperaturaCelsius)
			fmt.Println("La temperatura convertidad es ", Farenheit)
		} else if usrOpcion == "4" {
			fmt.Print("Ingresa una temperatura en Farenheit para convertir a Celsius: ")
			var temperaturaFarenheit float64
			fmt.Scan(&temperaturaFarenheit)
			Celsius := opcion4(temperaturaFarenheit)
			fmt.Println("La temperatura convertidad es ", Celsius)
		} else if usrOpcion == "0" || usrOpcion == "salir" {
			fmt.Println("Saliendo del programa")
			break
		} else {
			fmt.Println("Opcion Invalida")
		}
	}
}

func averageGrade(listaNotas []float64) float64 {
	var sumaNotas float64 = 0.0
	var promedioNotasCalculador float64
	for _, notaPorNota := range listaNotas {
		sumaNotas += notaPorNota
	}
	promedioNotasCalculador = sumaNotas / float64(len(listaNotas))
	return promedioNotasCalculador
}
func opcion1(numeroEstudiantes int) float64 {
	var nota float64
	var listaNotas []float64
	for contadorEstu := 1; contadorEstu <= numeroEstudiantes; contadorEstu++ {
		fmt.Print("Ingrese la nota del estudiante ", contadorEstu, " del 0 al 100: ")
		fmt.Scan(&nota)
		listaNotas = append(listaNotas, nota)
	}
	promedioNotas := averageGrade(listaNotas)
	if promedioNotas >= 70 {
		fmt.Println("Aprobado")
	} else {
		fmt.Println("Reprobado")
	}
	switch {
	case promedioNotas >= 70 && promedioNotas <= 79:
		fmt.Println("Satisfactory performance")
	case promedioNotas >= 80 && promedioNotas <= 89:
		fmt.Println("Good performance")
	case promedioNotas >= 90 && promedioNotas <= 100:
		fmt.Println("Excellent performance")
	case promedioNotas < 70:
		fmt.Println("Needs improvement")
	}
	return promedioNotas
}
func opcion2(n int) int {
	var totalSu int
	for contadorsuma := 1; contadorsuma <= n; contadorsuma++ {
		totalSu += contadorsuma
	}
	return totalSu
}
func opcion3(temperaturaCelsius float64) float64 {
	var Farenheit float64
	Farenheit = ((temperaturaCelsius * 9 / 5) + 32)
	return Farenheit
}
func opcion4(temperaturaFarenheit float64) float64 {
	var Celsius float64 = ((temperaturaFarenheit - 32) * 5 / 9)
	return Celsius
}
