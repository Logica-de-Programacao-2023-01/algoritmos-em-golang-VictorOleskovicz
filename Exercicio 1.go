package main

import "fmt"

func main() {
	var x1, x2, x3 float64
	fmt.Print(" Qual é x1? ")
	fmt.Scan(&x1)
	fmt.Print(" Qual é x2? ")
	fmt.Scan(&x2)
	fmt.Print(" Qual é x3? ")
	fmt.Scan(&x3)

	soma := x1 + x2 + x3

	fmt.Print(" A soma é ", soma)
}
