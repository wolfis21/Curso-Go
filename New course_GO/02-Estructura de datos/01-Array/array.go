package main

import "fmt"

func main() {
	//array

	var numeros [5]int

	numeros[0] = 1
	numeros[1] = 2
	numeros[2] = 3
	numeros[3] = 4
	numeros[4] = 5

	fmt.Println(numeros)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
	//arrays con valores
	nombres := [3]string{"Isaac", "Mariam", "alejandro"}

	//sin necesidad de hacer un bucle imprime el array
	fmt.Println(nombres)

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
	//tipo 2 definir array sin indice
	colores := [...]string{
		"Rojo",
		"Verde",
		"Azul",
	}
	//len es para saber la cantidad de indices del array
	fmt.Println(colores, len(colores))
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")

	//indice definido
	monedas := [...]string{
		0: "Dolar",
		2: "Euro",
		3: "bs",
		5: "pesos",
	}
	//como se salta el indice 1 aun se sigue contando que el array es de 6
	fmt.Println(monedas, len(monedas))

	//sub array
	submoneda := monedas[3:]
	fmt.Println(submoneda)

	//consiste en crear un arrego del otro, se espefica e inicio y el final

}
