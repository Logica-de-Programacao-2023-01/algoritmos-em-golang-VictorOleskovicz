package main

import "fmt"

func main() {
	var num1, num2, num3 int

	fmt.Println("Digite três números inteiros:")
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	fmt.Scanln(&num3)

	if num1 <= num2 && num1 <= num3 {
		fmt.Println("O menor número é:", num1)
	} else if num2 <= num1 && num2 <= num3 {
		fmt.Println("O menor número é:", num2)
	} else {
		fmt.Println("O menor número é:", num3)
	}
}
