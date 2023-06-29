package main

import "fmt"

func main() {
	var salario float64

	fmt.Println("Digite o salário:")
	fmt.Scanln(&salario)

	aumento := salario * 0.15
	novoSalario := salario + aumento

	fmt.Println("O novo salário é:", novoSalario)
}
