package main

import "fmt"

func main() {
	var x int
	fmt.Print("Qual é x")
	fmt.Scan(&x)

	dobro := x * 2
	triplo := x * 3
	Quadruplo := x * 4

	fmt.Print(" O dobro é ", dobro)
	fmt.Print(" O triplo é ", triplo)
	fmt.Print(" O Quadruplo é ", Quadruplo)
}
