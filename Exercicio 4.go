package main

import "fmt"

func main() {
	var x1, x2, x3 int
	fmt.Print(" Qual é x1? ")
	fmt.Scan(&x1)
	fmt.Print(" Qual é x2? ")
	fmt.Scan(&x2)
	fmt.Print(" Qual é x3? ")
	fmt.Scan(&x3)

	média_ponderada := ((x1 * 2) + (x2 * 3) + (x3 * 5)) / 10

	fmt.Print(" A média é ", média_ponderada)
}
