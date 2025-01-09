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

func Ejercicio4(calificaciones ...int) (min int, max int, prom int) {
	funcMinimo, err1 := obtenerFuncionPorTipoCalculo("minimo")
	if err1 != nil {
		fmt.Println(err1.Error())
	} else {
		min = funcMinimo(calificaciones...)
	}

	funcMaximo, err2 := obtenerFuncionPorTipoCalculo("maximo")
	if err2 != nil {
		fmt.Println(err2.Error())
	} else {
		max = funcMaximo(calificaciones...)
	}

	funcPromedio, err3 := obtenerFuncionPorTipoCalculo("promedio")
	if err3 != nil {
		fmt.Println(err3.Error())
	} else {
		prom = funcPromedio(calificaciones...)
	}

	funcError, err4 := obtenerFuncionPorTipoCalculo("invalid")
	if err4 != nil {
		fmt.Println(err4.Error())
	} else {
		funcError(calificaciones...)
	}

	return
}

func obtenerFuncionPorTipoCalculo(tipoCalculo string) (func(calificaciones ...int) int, error) {
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

func calcularMinimoCalificaciones(calificaciones ...int) int {
	valorMinimo := calificaciones[0]
	for _, calificacion := range calificaciones {
		if calificacion < valorMinimo {
			valorMinimo = calificacion
		}
	}
	fmt.Printf("El valor mínimo de las calificaciones es: %d", valorMinimo)
	return valorMinimo
}

func calcularMaximoCalificaciones(calificaciones ...int) int {
	valorMaximo := calificaciones[0]
	for _, calificacion := range calificaciones {
		if calificacion > valorMaximo {
			valorMaximo = calificacion
		}
	}
	fmt.Printf("El valor máximo de las calificaciones es: %d", valorMaximo)
	return valorMaximo
}

func calcularPromedioCalificaciones(calificaciones ...int) int {
	sumaCalificaciones := 0
	for _, calificacion := range calificaciones {
		sumaCalificaciones += calificacion
	}
	promedio := sumaCalificaciones / len(calificaciones)

	fmt.Printf("El promedio de las calificaciones es: %d", promedio)
	return promedio
}
