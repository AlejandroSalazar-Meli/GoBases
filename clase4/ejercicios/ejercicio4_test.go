package ejercicios

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEjercicioCuatro(t *testing.T) {
	calificaciones := []int{1, 2, 3, 4, 5}
	minimoEsperado := 1
	maximoEsperado := 5
	promedioEsperado := 3

	minimoObtenido, maximoObtenido, promedioObtenido := Ejercicio4(calificaciones...)

	require.Equal(t, minimoEsperado, minimoObtenido)
	require.Equal(t, maximoEsperado, maximoObtenido)
	require.Equal(t, promedioEsperado, promedioObtenido)

}
