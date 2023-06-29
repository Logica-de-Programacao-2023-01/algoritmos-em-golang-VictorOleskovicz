package main

import "fmt"

func main() {
	var num int

	fmt.Println("Digite um número inteiro:")
	fmt.Scanln(&num)

	sucessor := num + 1
	antecessor := num - 1

	fmt.Println("O sucessor é:", sucessor)
	fmt.Println("O antecessor é:", antecessor)
}
