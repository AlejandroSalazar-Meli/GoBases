package ejercicios

import (
	"errors"
	"fmt"
)

const (
	minimo   = "minimo"
	maximo   = "maximo"
	promedio = "promedio"
)

func Ejercicio4() {
	funcMinimo, err1 := obtenerFuncionPorTipoCalculo("minimo")
	if err1 != nil {
		fmt.Println(err1.Error())
	} else {
		funcMinimo(2, 3, 4, 5)
	}

	funcMaximo, err2 := obtenerFuncionPorTipoCalculo("maximo")
	if err2 != nil {
		fmt.Println(err2.Error())
	} else {
		funcMaximo(1, 2, 3, 4, 5)
	}

	funcPromedio, err3 := obtenerFuncionPorTipoCalculo("promedio")
	if err3 != nil {
		fmt.Println(err3.Error())
	} else {
		funcPromedio(1, 2, 3, 4, 5)
	}

	funcError, err4 := obtenerFuncionPorTipoCalculo("invalid")
	if err4 != nil {
		fmt.Println(err4.Error())
	} else {
		funcError(2, 3, 4, 5)
	}
}

func obtenerFuncionPorTipoCalculo(tipoCalculo string) (func(calificaciones ...int), error) {
	switch tipoCalculo {
	case minimo:
		return calcularMinimoCalificaciones, nil
	case maximo:
		return calcularMaximoCalificaciones, nil
	case promedio:
		return calcularPromedioCalificaciones, nil
	default:
		return nil, errors.New("operacion invalida")
	}
}

func calcularMinimoCalificaciones(calificaciones ...int) {
	valorMinimo := calificaciones[0]
	for _, calificacion := range calificaciones {
		if calificacion < valorMinimo {
			valorMinimo = calificacion
		}
	}
	fmt.Printf("El valor mínimo de las calificaciones es: %d", valorMinimo)
}

func calcularMaximoCalificaciones(calificaciones ...int) {
	valorMaximo := calificaciones[0]
	for _, calificacion := range calificaciones {
		if calificacion > valorMaximo {
			valorMaximo = calificacion
		}
	}
	fmt.Printf("El valor máximo de las calificaciones es: %d", valorMaximo)
}

func calcularPromedioCalificaciones(calificaciones ...int) {
	sumaCalificaciones := 0.0
	for _, calificacion := range calificaciones {
		sumaCalificaciones += float64(calificacion)
	}
	promedio := sumaCalificaciones / float64(len(calificaciones))

	fmt.Printf("El promedio de las calificaciones es: %.2f", promedio)
}
