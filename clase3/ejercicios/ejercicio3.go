package ejercicios

import (
	"fmt"
)

func Ejercicio3(categoria string, minutosTrabajados int) {

	switch categoria {
	case "C":
		imprimirRespuesta(calcularCantidadHoras(minutosTrabajados) * 1000)
	case "B":
		imprimirRespuesta((calcularCantidadHoras(minutosTrabajados) * 1500) * 1.2)
	case "A":
		imprimirRespuesta((calcularCantidadHoras(minutosTrabajados) * 3000) * 1.5)
	default:
		fmt.Printf("categoría inválida (A, B, C)")
	}

}

func calcularCantidadHoras(cantidadMinutos int) float64 {
	return float64(cantidadMinutos) / 60
}

func imprimirRespuesta(salario float64) {
	fmt.Printf("El salario del empleado es: %f", salario)
}
