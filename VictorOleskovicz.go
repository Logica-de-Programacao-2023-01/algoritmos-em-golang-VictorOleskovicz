package main

import "fmt"

func main() {
	var cambio, dolares float64
	fmt.Print(" Qual o valor do dólar  em reais?")
	fmt.Scan(&cambio)
	fmt.Print(" Quantos reais você quer converter")
	fmt.Scan(&dolares)

	reais := dolares * cambio

	fmt.Println(" O valor em reais é de", reais)
}
