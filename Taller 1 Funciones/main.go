package main

import "fmt"

func main() {
	var usrOpcion string
	fmt.Println("Ingresa tu opcion del 1 al 4 o 0 para salir ")
	fmt.Println("La opcion 1 es para calcular la nota de estudiantes de un curso del 1 al 100")
	fmt.Println("La opcion 2 es para calcular la suma desde 1 al numero elegido")
	fmt.Println("La opcion 3 para convertir de Celsius a Farenheit")
	fmt.Println("La opcion 4 es para convertir de Farenheit a Celsius")
	fmt.Println("La opcion 0 o salir termina el programa")
	fmt.Scan(&usrOpcion)
	switch usrOpcion {
	case usrOpcion == 1:
		fmt.Println("Ingrese el numero de estudiantes")
		var numeroEstudiantes int
		fmt.Scan(&numeroEstudiantes)
	case usrOpcion == 2:
		fmt.Println("Ingresa un numero para calcular la suma")
		var n int
		fmt.Scan(&n)
		fmt.Println("La suma de los numeros: ", totalSu)
	case usrOpcion == 3:
		fmt.Println("Ingresa una temperatura en Celsius para convertir a Farenheit")
		var temperaturaCelsius float64
		fmt.Scan(&temperaturaCelsius)
		fmt.Println("La temperatura convertidad es ", opcion3(Farenheit))
	case usrOpcion == 4:
		fmt.Println("Ingresa una temperatura en Farenheit para convertir a Celsius")
		var temperaturaFarenheit float64
		fmt.Scan(&temperaturaFarenheit)
		fmt.Println("La temperatura convertidad es ", opcion4(Celsius))
	case usrOpcion == 0 && usrOpcion == "salir":
		break
	default:
		fmt.Println("Opcion Invalida")
	}
}

func opcion(numeroEstudiantes int) float64 {
	nota := []float64
	var sumaNotas float64 = 0.0
	for contadorEstu := 1; contadorEstu <= numeroEstudiantes; contadorEstu++ {
		fmt.Println("Ingrese la nota del estudiante del 0 al 100", contadorEstu, ":")
		fmt.Scan(&nota)
		for _, nota := range nota {
			sumaNotas += nota
		}
	}
	var promedioNotas float64 = sumaNotas / float64(len(nota))

	if promedioNotas >= 70 {
		fmt.Println("Aprobado")
	} else {
		fmt.Println("Reprobado")
	}
	switch promedioNotas {
	case promedioNotas >= 70.0 && promedioNotas <= 79.99:
		fmt.Println("Satisfactory performance")
	case promedioNotas >= 80.0 && promedioNotas <= 89.99:
		fmt.Println("Good performance")
	case promedioNotas >= 90.0 && promedioNotas <= 100:
		fmt.Println("Good performance")
	case promedioNotas < 70:
		fmt.Println("Needs improvement")
	}

}
func opcion2(n int) int {
	totalSu := 0
	for contador := 1; contador <= n; contador++ {
		var nota int
		nota += totalSu
	}
	return totalSu
}
func opcion3(temperaturaCelsius float64) float64 {
	var Farenheit float64 = ((temperaturaCelsius * (9 / 5)) + 32)
	return Farenheit
}
func opcion4(temperaturaFarenheit float64) float64 {
	var Celsius float64 = ((temperaturaFarenheit - 32) * 5 / 9)
	return Celsius
}
