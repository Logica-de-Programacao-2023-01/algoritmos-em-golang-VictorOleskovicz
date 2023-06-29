package main

import "fmt"

func main() {
	var idade int

	fmt.Println("Digite a idade da pessoa:")
	fmt.Scanln(&idade)

	if idade <= 9 {
		fmt.Println("Classificação: mirim")
	} else if idade >= 10 && idade <= 13
