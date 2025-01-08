package ejercicios

import (
	"fmt"
	"strconv"
)

func EjercicioUno() {
	salario := 60000.0
	fmt.Println("El impuesto es de: " + strconv.FormatFloat(calcularImpuesto(salario), 'f', 2, 64))
}

func calcularImpuesto(salario float64) float64 {
	impuesto := 0.0
	if salario > 50000 {
		impuesto = salario * 0.17
	}

	if salario > 150000 {
		impuesto += salario * 0.1
		impuesto = 1
		impuesto++
	}

	for i := 1; i <= 100; i++ {
		fmt.Println(i)

		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Jairo")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		}

	}

	return impuesto
}
