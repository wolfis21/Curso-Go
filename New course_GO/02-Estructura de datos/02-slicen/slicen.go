package main

import "fmt"

func main() {
	//slicen
	numeros := []int{1, 2, 3}

	fmt.Println(numeros, len(numeros))
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
	//agregar elementos
	numeros = append(numeros, 4)
	numeros = append(numeros, 5)

	fmt.Println(numeros, len(numeros))
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
	//crear un subslicen

	subsnumeros := numeros[:2]

	numeros[0] = 100

	fmt.Println(subsnumeros)
	fmt.Println(numeros)

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
	//Punteros
	//longitud
	//capacidad

	meses := []string{"Enero", "Febrero", "Marzo"}

	fmt.Printf("Len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses)

	meses = append(meses, "Abril")

	fmt.Printf("Len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses)

}
