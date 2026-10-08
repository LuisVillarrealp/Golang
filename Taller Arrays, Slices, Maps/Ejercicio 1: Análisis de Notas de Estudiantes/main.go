package main

import "fmt"

func main() {
	var notas [6][4]float64
	var promedios []float64

	fmt.Println("Analisis de Notas de Estudiantes")
	for contadorEstudiante := 0; contadorEstudiante < 6; contadorEstudiante++ {
		for contadorNotaEstudiante := 0; contadorNotaEstudiante < 4; contadorNotaEstudiante++ {
			fmt.Println("Ingresa la nota del estudiante", contadorEstudiante+1, "en la materia", contadorNotaEstudiante+1, ": ")
			fmt.Scan(&notas[contadorEstudiante][contadorNotaEstudiante])
		}
	}

	fmt.Println("			RESULTADOS			")
	for arrayEstudiante := 0; arrayEstudiante < 6; arrayEstudiante++ {
		var fila []float64 = notas[arrayEstudiante][:]
		prom := promedio(fila)
		mayor, menor := mayorMenornota(fila)
		promedios = append(promedios, prom)
		fmt.Println("	Estudiante ", arrayEstudiante+1)
		fmt.Println("Notas: ", fila)
		fmt.Println("Promedio: ", prom)
		fmt.Println("Nota mas alta: ", mayor)
		fmt.Println("Nota mas baja: ", menor)
	}
	fmt.Println("			Promedio			")
	fmt.Println("El promedio general de la clase es: ", promedio(promedios))
}

func promedio(notas []float64) float64 {
	var suma float64 = 0
	for _, nota := range notas {
		suma += nota
	}
	return suma / float64(len(notas))
}

func mayorMenornota(notas []float64) (float64, float64) {
	var mayor float64 = notas[0]
	var menor float64 = notas[0]
	for _, nota := range notas {
		if nota > mayor {
			mayor = nota
		}
		if nota < menor {
			menor = nota
		}
	}
	return mayor, menor
}
