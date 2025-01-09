package ejercicios

import "fmt"

func EjercicioDos(calificaciones ...int) int {
	fmt.Printf("El promedio es: %d\n", calcularPromedio(calificaciones...))
	return calcularPromedio(calificaciones...)
}

func calcularPromedio(valores ...int) int {
	suma := 0
	for _, valor := range valores {
		suma += valor
	}

	promedio := suma / len(valores)
	return promedio
}
