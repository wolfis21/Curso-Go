package main

import "fmt"

func saludar(nombre string) {
	// fmt.Println("Buen dia")
	fmt.Println("Buen dia", nombre)
}

//para las funciones con retorno se debe de especificar la salida
func sumar(a, b int) int {
	return a + b
}

func main() {
	nombre := "alejandro"

	saludar("Isaac")
	saludar(nombre)

	res := sumar(10, 15)

	fmt.Println(res)
}
