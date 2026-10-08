package main

import "fmt"

func main() {
	var votos map[string]int = map[string]int{
		"deportes":    0,
		"videojuegos": 0,
		"cine":        0,
		"musica":      0,
	}
	var usrVoto string

	fmt.Println("Bienvenido a la votacion de actividades de integracion")
	fmt.Println("Tus opciones son las siguiente: ")
	for contadorVotos := 0; contadorVotos < 5; {
		fmt.Println("			Votante numero", contadorVotos+1)
		fmt.Println("1. Deportes")
		fmt.Println("2. Videojuegos")
		fmt.Println("3. Cine")
		fmt.Println("4. Musica")
		fmt.Printf("Ingresa tu opcion: ")
		fmt.Scan(&usrVoto)
		switch usrVoto {
		case "1":
			votos["deportes"]++
			contadorVotos++
		case "2":
			votos["videojuegos"]++
			contadorVotos++
		case "3":
			votos["cine"]++
			contadorVotos++
		case "4":
			votos["musica"]++
			contadorVotos++
		default:
			fmt.Println("Opcion invalida, intenta de nuevo")
		}
	}

	fmt.Println("			RESULTADOS")
	for actividad, cantidadVotos := range votos {
		fmt.Println("La actividad ", actividad, " tiene ", cantidadVotos, " votos")
	}
	ganadora, votosGanadora := actividadMasVotada(votos)
	fmt.Println("La actividad con mas votos es: ", ganadora, " con ", votosGanadora, " votos")
}

func actividadMasVotada(votos map[string]int) (string, int) {
	var ganadora string
	var mayorVotos int = 0
	for actividad, cantidadVotos := range votos {
		if cantidadVotos > mayorVotos {
			ganadora = actividad
			mayorVotos = cantidadVotos
		}
	}
	return ganadora, mayorVotos
}
