package contadorvocales

func ContarVocales(palabra string) (int, int, int, int, int) {
	var a int = 0
	var e int = 0
	var i int = 0
	var o int = 0
	var u int = 0
	for _, letra := range palabra {
		switch letra {
		case 'a', 'A':
			a++
		case 'e', 'E':
			e++
		case 'i', 'I':
			i++
		case 'o', 'O':
			o++
		case 'u', 'U':
			u++
		}
	}
	return a, e, i, o, u
}
