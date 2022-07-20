package main

import "fmt"

func main() {

	hola := "Hola"
	mundo := "Mundo"

	// tipo 1 sin  salto de linea
	fmt.Print(hola, mundo)

	fmt.Print("\n")
	// tipo 2
	nombre := "Isaac"
	edad := 21
	fmt.Printf("Hola, %s y su edad es %d \n", nombre, edad)

	//tipo 2.1
	fmt.Printf("Hola, %v y su edad es %v \n", nombre, edad)

	// tipo 3 con salto de linea
	fmt.Println(hola, mundo)

	//forma de almacenar un mensaje en un variable

	mensaje := fmt.Sprintf("Holaaaa, %v y su edad es %v \n", nombre, edad)
	fmt.Println(mensaje)

	//ingreso de datos por teclado
	fmt.Printf("Ingrese un nombre: ")
	fmt.Scanln(&nombre)

	fmt.Println("ingresado es: ", nombre)

}
