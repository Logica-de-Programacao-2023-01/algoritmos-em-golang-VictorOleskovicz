package main

import "fmt"

func main() {
	var num, sum, count int
	fmt.Println("Digite os números (digite 0 para encerrar):")

	for {
		fmt.Print("Número: ")
		fmt.Scan(&num)

		if num == 0 {
			break
		}

		sum += num
		count++
	}

	if count > 0 {
		average := float64(sum) / float64(count)
		fmt.Printf("Média: %.2f\n", average)
	} else {
		fmt.Println("Nenhum número foi inserido.")
	}
}
