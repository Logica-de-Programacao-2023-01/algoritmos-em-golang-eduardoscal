package main

import "fmt"

func main() {
	var numero int

	fmt.Println("Digite um número inteiro positivo:")
	fmt.Scanln(&numero)

	fmt.Println("Divisores de", numero, ":")

	for i := 1; i <= numero; i++ {
		if numero%i == 0 {
			fmt.Println(i)
		}
	}
}
