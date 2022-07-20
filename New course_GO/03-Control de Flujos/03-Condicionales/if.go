package main

import "fmt"

func main() {

	var consumo, descuento float64
	var datosdescuento string

	fmt.Print("Ingrese el consumo: ")
	fmt.Scanln(&consumo)

	if consumo > 0 {
		if consumo <= 100 {
			datosdescuento = "10%"
			descuento = consumo * 0.10

		} else if consumo > 100 && consumo <= 200 {

			datosdescuento = "20%"
			descuento = consumo * 0.20

		} else if consumo > 200 {
			datosdescuento = "30%"
			descuento = consumo * 0.30

		}
	} else {
		fmt.Println("dato erroneo")
	}

	montoDescuento := consumo - descuento
	igv := montoDescuento * 0.19

	total := montoDescuento + igv

	fmt.Println("------------------------------")
	fmt.Println("FACTURA")
	fmt.Println("------------------------------")
	fmt.Println("Descuento a aplicar: ", datosdescuento)
	fmt.Println("Consumo: ", consumo)
	fmt.Println("Descuento: ", descuento)
	fmt.Println("Monto con descuento: ", montoDescuento)
	fmt.Println("IGV: ", igv)
	fmt.Println("Total a pagar: ", total)

}
