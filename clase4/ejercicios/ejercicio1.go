package ejercicios

import (
	"fmt"
)

func EjercicioUno(salario float64) float64 {
	fmt.Printf("El impuesto es de: %2f\n", calcularImpuesto(salario))
	return calcularImpuesto(salario)
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
