package ejercicios

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPerro(t *testing.T) {
	cantidadAnimal := 10
	cantidadAlimentoEsperado := 100.0

	cantidadAlimentoObtenida, _, _, _ := Ejercicio5(cantidadAnimal)

	require.Equal(t, cantidadAlimentoEsperado, cantidadAlimentoObtenida)
}

func TestGato(t *testing.T) {
	cantidadAnimal := 10
	cantidadAlimentoEsperado := 50.0

	_, cantidadAlimentoObtenida, _, _ := Ejercicio5(cantidadAnimal)

	require.Equal(t, cantidadAlimentoEsperado, cantidadAlimentoObtenida)
}

func TestHamster(t *testing.T) {
	cantidadAnimal := 10
	cantidadAlimentoEsperado := 2.5

	_, _, cantidadAlimentoObtenida, _ := Ejercicio5(cantidadAnimal)

	require.Equal(t, cantidadAlimentoEsperado, cantidadAlimentoObtenida)
}

func TestTarantula(t *testing.T) {
	cantidadAnimal := 10
	cantidadAlimentoEsperado := 1.5

	_, _, _, cantidadAlimentoObtenida := Ejercicio5(cantidadAnimal)

	require.Equal(t, cantidadAlimentoEsperado, cantidadAlimentoObtenida)
}
