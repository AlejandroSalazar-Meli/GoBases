package ejercicios

import "fmt"

func EjercicioDos() {
	fmt.Printf("El promedio es: %d", calcularPromedio(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

func calcularPromedio(valores ...int) int {
	suma := 0
	for _, valor := range valores {
		suma += valor
	}

	promedio := suma / len(valores)
	return promedio
}
