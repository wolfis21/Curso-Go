package main

import "fmt"

func cociente(a, b int) int {
	return a / b
}
func residuo(a, b int) int {
	return a % b
}

func main() {

	var a int
	var b int

	fmt.Printf("Escriba un numero: ")
	fmt.Scanln(&a)

	fmt.Printf("Escriba otro: ")
	fmt.Scanln(&b)

	resp := cociente(a, b)
	resp2 := residuo(a, b)

	fmt.Println("\nEl cociente es: ", resp)
	fmt.Println("\nEl residuo es: ", resp2)
}
