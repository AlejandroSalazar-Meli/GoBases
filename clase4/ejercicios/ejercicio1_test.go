package ejercicios

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEjercicioUnoMenosDe50000(t *testing.T) {
	salario1 := 10000.0
	salario2 := 100000.0
	salario3 := 200000.0

	impuestoEsperado1 := 0.0
	impuestoEsperado2 := 17000.0
	impuestoEsperado3 := 54000.0

	impuestoObtenido1 := EjercicioUno(salario1)
	impuestoObtenido2 := EjercicioUno(salario2)
	impuestoObtenido3 := EjercicioUno(salario3)

	require.Equal(t, impuestoEsperado1, impuestoObtenido1)
	require.Equal(t, impuestoEsperado2, impuestoObtenido2)
	require.Equal(t, impuestoEsperado3, impuestoObtenido3)
}
