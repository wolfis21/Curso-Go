package main

import "fmt"

func main() {
	var vocal string

	fmt.Print("Ingrese una letra: ")
	fmt.Scanln(&vocal)

	//metodo 1
	// switch {
	// case vocal == "a" || vocal == "A":
	// 	fmt.Println(vocal, "Es vocal")

	// case vocal == "e" || vocal == "E":
	// 	fmt.Println(vocal, "Es vocal")

	// case vocal == "i" || vocal == "I":
	// 	fmt.Println(vocal, "Es vocal")

	// case vocal == "o" || vocal == "O":
	// 	fmt.Println(vocal, "Es vocal")

	// case vocal == "u" || vocal == "U":
	// 	fmt.Println(vocal, "Es vocal")

	// default:
	// 	fmt.Println(vocal, "No es un Vocal")
	// }

	//metodo 2

	switch vocal {
	case "a", "A", "e", "E", "i", "I", "o", "O", "u", "U":
		fmt.Println(vocal, "Es Vocal")
	default:
		fmt.Println(vocal, "No es un Vocal")
	}
}
