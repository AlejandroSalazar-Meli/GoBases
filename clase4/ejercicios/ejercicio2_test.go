package ejercicios

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEjercicioDos(t *testing.T) {
	calificaciones := []int{1, 2, 3, 4, 5}
	promedioEsperado := 3
	promedioObtenido := EjercicioDos(calificaciones...)

	require.Equal(t, promedioEsperado, promedioObtenido)
}
