package main

import (
	"fmt"
)

func main() {
	dias := make(map[int]string)

	fmt.Println(dias)

	//agregar datos
	dias[1] = "Domingo"
	dias[2] = "Lunes"
	dias[3] = "Martes"
	dias[4] = "Miercoles"
	dias[5] = "Jueves"
	dias[6] = "Viernes"
	dias[7] = "Sabado"

	fmt.Println(dias)

	//puedes modificar valor ya colocado
	dias[7] = "SABADO"

	fmt.Println(dias)

	//se puede eleminar elemento del array completo
	delete(dias, 1)

	fmt.Println(dias)

	//Nueva mapa
	estudiantes := make(map[string][]int)

	estudiantes["Isaac"] = []int{17, 18, 19}
	estudiantes["Mariam"] = []int{19, 19, 17}

	fmt.Println(estudiantes)
	//para mostrar un elemento en concreto
	fmt.Println(estudiantes["Mariam"][1])
}
