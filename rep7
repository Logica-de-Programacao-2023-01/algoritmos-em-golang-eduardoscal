package main

import "fmt"

func main() {
	var salario float64

	fmt.Println("Digite o salário do funcionário:")
	fmt.Scanln(&salario)

	var novoSalario float64

	if salario < 1000.00 {
		novoSalario = salario * 1.10
	} else {
		novoSalario = salario * 1.05
	}

	fmt.Printf("O novo salário é: %.2f\n", novoSalario)
}
