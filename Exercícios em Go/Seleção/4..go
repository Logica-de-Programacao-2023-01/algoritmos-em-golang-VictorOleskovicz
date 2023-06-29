package main

import "fmt"

func main() {
	var altura float64
	var sexo string

	fmt.Println("Digite a altura (em metros):")
	fmt.Scanln(&altura)

	fmt.Println("Digite o sexo (M para masculino ou F para feminino):")
	fmt.Scanln(&sexo)

	pesoIdeal := 0.0

	if sexo == "M" {
		pesoIdeal = 72.7*altura - 58
	} else if sexo == "F" {
		pesoIdeal = 62.1*altura - 44.7
	}

	fmt.Println("O peso ideal para essa pessoa é:", pesoIdeal)

	var peso float64

	fmt.Println("Digite o peso atual (em kg):")
	fmt.Scanln(&peso)

	if peso < pesoIdeal {
		fmt.Println("A pessoa está abaixo do peso ideal.")
	} else if peso > pesoIdeal {
		fmt.Println("A pessoa está acima do peso ideal.")
	} else {
		fmt.Println("A pessoa está dentro do peso ideal.")
	}
}
