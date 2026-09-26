package contadorvocales

import (
	"bufio"
	"fmt"
	"os"
)

func ContarVocales(frase string) (int, int, int, int, int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Printf("Tu Frase aqui: ")
	scanner.Scan()
	frase = scanner.Text()

	var a int = 0
	var e int = 0
	var i int = 0
	var o int = 0
	var u int = 0
	for _, letra := range frase {
		switch letra {
		case 'a', 'A', 'á', 'Á':
			a++
		case 'e', 'E', 'é', 'É':
			e++
		case 'i', 'I', 'í', 'Í':
			i++
		case 'o', 'O', 'ó', 'Ó':
			o++
		case 'u', 'U', 'ú', 'Ú', 'ü', 'Ü':
			u++
		}
	}
	return a, e, i, o, u
}
