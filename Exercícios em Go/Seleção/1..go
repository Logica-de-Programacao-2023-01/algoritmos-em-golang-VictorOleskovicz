package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Println("Digite dois números inteiros:")
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)

	if num1 > num2 {
		fmt.Println("O maior número é:", num1)
	} else {
		fmt.Println("O maior número é:", num2)
	}
}
