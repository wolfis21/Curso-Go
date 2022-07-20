package main

import "fmt"

func main() {
	var (
		nombre string
		edad   int
		pi     float64
	)
	fmt.Println("Ingrese su Nombre: ")
	fmt.Scanln(&nombre)

	fmt.Println("Ingrese edad: ")
	fmt.Scanln(&edad)

	pi = 3.1416

	fmt.Printf("Nombre: %s Edad: %d \nValor de pi: %f", nombre, edad, pi)
}
