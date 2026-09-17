package main

import "fmt"

func main() {
	var n int
	fmt.Println("Ingresa un numero entero positivo")
	fmt.Scanf("%d", &n)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, n*i)
	}

	if n%2 == 0 {
		fmt.Println("El número es Par")
	} else {
		fmt.Println("El número es Impar")
	}

	switch {
	case n >= 1 && n <= 5:
		fmt.Println("Número pequeño")
	case n >= 6 && n <= 10:
		fmt.Println("Número mediano")
	case n > 10:
		fmt.Println("Número grande")
	}
}
