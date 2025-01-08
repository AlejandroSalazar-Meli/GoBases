package ejercicioclase

import "fmt"

func EjercicioUno() {
	var palabra string
	fmt.Print("Ingresa la palabra: ")
	_, err := fmt.Scan(&palabra)
	if err != nil {
		fmt.Println("Error en el input: ", err)
	} else {
		fmt.Println("La cantidad de letras en la palabra es: ", len(palabra))
		for _, letra := range palabra {
			fmt.Println(string(letra))
		}
	}
}

func EjercicioDos() {

}
