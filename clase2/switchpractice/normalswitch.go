package switchpractice

import (
	"fmt"
)

func NormalSwitch() {
	var paisNacimiento string
	fmt.Print("Ingresa tu país de nacimiento: ")
	_, err := fmt.Scan(&paisNacimiento)
	if err != nil {
		fmt.Println("Error en el input: ", err)
	} else {
		switch paisNacimiento {
		case "colombia", "Colombia":
			fmt.Println("Parcero!")
		case "argentina", "Argentina":
			fmt.Println("Che!")
		case "mexico", "Mexico":
			fmt.Println("Wey!")
		default:
			fmt.Println("Amigo!")
		}
	}
}
