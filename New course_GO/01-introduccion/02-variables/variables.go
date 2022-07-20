package main

import (
	"fmt"
)

func main() {
	//formas de declarar
	// metodo 1
	var nombre1 string
	nombre1 = "Isaac"

	// metodo 2
	var nombre2 string = "Alejandro"

	// metodo 3
	edad := 21

	fmt.Println(nombre1, nombre2, edad)

	//TIPOS DE VARIABLES
	var a int
	var b float64
	var b2 float32
	var c string
	var d bool
	//si no se declaran sus valores por defecto son
	fmt.Println(a, b, b2, c, d)

	//valores constantes
	const pi = 3.14

	fmt.Println(pi)
}
