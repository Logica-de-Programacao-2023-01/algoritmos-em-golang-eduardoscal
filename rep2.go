package main

import "fmt"

func main() {
	var num1, num2, num3 int

	fmt.Println("Digite o primeiro número:")
	fmt.Scanln(&num1)
	fmt.Println("Digite o segundo número:")
	fmt.Scanln(&num2)
	fmt.Println("Digite o terceiro número:")
	fmt.Scanln(&num3)

	if num1 < num2 && num1 < num3 {
		fmt.Println("O menor número é:", num1)
	} else if num2 < num1 && num2 < num3 {
		fmt.Println("O menor número é:", num2)
	} else if num3 < num1 && num3 < num2 {
		fmt.Println("O menor número é:", num3)
	} else {
		fmt.Println("Os números são iguais.")
	}
}
