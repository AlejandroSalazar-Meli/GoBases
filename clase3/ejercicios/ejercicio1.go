package ejercicios

import (
	"fmt"
)

func EjercicioUno() {
	salario := 160000.0
	fmt.Printf("El impuesto es de: %2f", calcularImpuesto(salario))
}

func calcularImpuesto(salario float64) float64 {
	impuesto := 0.0
	if salario > 50000 {
		impuesto = salario * 0.17
	}

	if salario > 150000 {
		impuesto += salario * 0.1
	}

	return impuesto
}
