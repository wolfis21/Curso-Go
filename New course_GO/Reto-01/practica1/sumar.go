package main

import (
	"fmt"
)

func sumar(a, b float64) float64 {
	return a + b
}
func main() {

	var a float64
	var b float64

	fmt.Printf("Escriba un numero: ")
	fmt.Scanln(&a)

	fmt.Printf("Escriba otro: ")
	fmt.Scanln(&b)

	resp := sumar(a, b)

	fmt.Println("\nEl resultado es: ", resp)
}
