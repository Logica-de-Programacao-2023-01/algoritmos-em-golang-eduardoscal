package main

import "fmt"

func main() {
	var altura, pesoIdeal float64
	var sexo string

	fmt.Println("Digite a altura (em metros):")
	fmt.Scanln(&altura)
	fmt.Println("Digite o sexo (M ou F):")
	fmt.Scanln(&sexo)

	if sexo == "M" {
		pesoIdeal = (72.7 * altura) - 58
	} else if sexo == "F" {
		pesoIdeal = (62.1 * altura) - 44.7
	}

	fmt.Println("O peso ideal é:", pesoIdeal)

	var peso float64

	fmt.Println("Digite o peso atual:")
	fmt.Scanln(&peso)

	if peso < pesoIdeal {
		fmt.Println("Abaixo do peso ideal.")
	} else if peso == pesoIdeal {
		fmt.Println("Dentro do peso ideal.")
	} else {
		fmt.Println("Acima do peso ideal.")
	}
}
