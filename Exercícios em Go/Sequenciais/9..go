package main

import "fmt"

func main() {
	var preco float64

	fmt.Println("Digite o preço do produto:")
	fmt.Scanln(&preco)

	desconto := preco * 0.1
	precoComDesconto := preco - desconto

	fmt.Println("O valor do produto com desconto é:", precoComDesconto)
}
