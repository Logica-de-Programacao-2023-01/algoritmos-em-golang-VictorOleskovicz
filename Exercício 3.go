package main

import "fmt"

func main() {
	var peso, altura float64
	fmt.Print(" Qual seu peso? ")
	fmt.Scan(&peso)
	fmt.Print(" Qual sua altura? ")
	fmt.Scan(&altura)

	IMC := peso / (altura * altura)

	fmt.Print(" Seu IMC é ", IMC)
}
