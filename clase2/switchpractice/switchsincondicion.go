package switchpractice

import "fmt"

func SwitchSinCondicion() {
	var temperatura int
	fmt.Print("Ingresa la temperatura: ")
	_, err := fmt.Scan(&temperatura)
	if err != nil {
		fmt.Println("Error en el input: ", err)
	} else {
		switch {
		case temperatura <= 17:
			fmt.Println("que frío!!")
		case temperatura > 17 && temperatura <= 27:
			fmt.Println("A gustito")
			fallthrough
		case temperatura > 27:
			fmt.Println("Que calor!")
		}
	}
}
