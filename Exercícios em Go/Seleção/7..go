package main

import "fmt"

func main() {
	var salario float64

	fmt.Println("Digite o salário do funcionário:")
	fmt.Scanln(&salario)

	var aumento float64

	if salario < 1000.00 {
		aumento = salario * 0.10
	} else {
		aumento = salario * 0.05
	}

	novoSalario := salario + aumento

	fmt.Println("O novo salário do funcionário é:", novoSalario)
}
