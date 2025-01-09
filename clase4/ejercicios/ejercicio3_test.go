package ejercicios

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEjercicioTres(t *testing.T) {
	cantidadMinutosA := 600
	cantidadMinutosB := 600
	cantidadMinutosC := 600
	cantidadMinutosD := 600

	salarioEsperadoA := 45000.0
	salarioEsperadoB := 18000.0
	salarioEsperadoC := 10000.0
	salarioEsperadoD := 0.0

	salarioObtenidoA := Ejercicio3("A", cantidadMinutosA)
	salarioObtenidoB := Ejercicio3("B", cantidadMinutosB)
	salarioObtenidoC := Ejercicio3("C", cantidadMinutosC)
	salarioObtenidoD := Ejercicio3("D", cantidadMinutosD)

	require.Equal(t, salarioEsperadoA, salarioObtenidoA)
	require.Equal(t, salarioEsperadoB, salarioObtenidoB)
	require.Equal(t, salarioEsperadoC, salarioObtenidoC)
	require.Equal(t, salarioEsperadoD, salarioObtenidoD)
}
