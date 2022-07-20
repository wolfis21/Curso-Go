package main

import "fmt"

func igv(a float64) float64 {
	return a * 0.18
}
func operacion(a float64) float64 {
	r := igv(a)
	return a + r
}

func main() {

	var x float64 //precio base
	// var y float64 //precio de venta

	fmt.Printf("Ingrese precio base: ")
	fmt.Scanln(&x)

	resp := igv(x)
	resp2 := operacion(x)

	fmt.Println("==========================")
	fmt.Println("El precio base es:", x)
	fmt.Println("El igv es:", resp)
	fmt.Println("y con un precio total de:", resp2)
	fmt.Println("==========================")
}
