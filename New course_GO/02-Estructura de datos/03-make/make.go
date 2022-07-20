package main

import "fmt"

func main() {

	numeros := make([]int, 0, 3)

	numeros = append(numeros, 100)
	numeros = append(numeros, 200)
	numeros = append(numeros, 300)

	fmt.Println(numeros)
}
