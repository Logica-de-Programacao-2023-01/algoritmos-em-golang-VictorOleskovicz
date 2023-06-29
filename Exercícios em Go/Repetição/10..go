package main

import "fmt"

func main() {
	var num, max int
	fmt.Println("Digite os números (digite 0 para encerrar):")

	for {
		fmt.Print("Número: ")
		fmt.Scan(&num)

		if num == 0 {
			break
		}

		if num > max {
			max = num
		}
	}

	if max != 0 {
		fmt.Println("O maior número é:", max)
	} else {
		fmt.Println("Nenhum número foi inserido.")
	}
}
