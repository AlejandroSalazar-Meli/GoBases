package ejercicios

import "fmt"

type Alumno struct {
	Nombre       string
	Apellido     string
	Dni          string
	FechaIngreso string
}

func (alumno Alumno) detalle() {
	fmt.Printf("El alumno %s %s, con DNI %s y fecha de ingreso %s", alumno.Nombre, alumno.Apellido, alumno.Dni, alumno.FechaIngreso)
}

func EjercicioUno() {
	alumno := Alumno{
		Nombre:       "Alejandro",
		Apellido:     "Salazar",
		Dni:          "1036",
		FechaIngreso: "30-07-2010",
	}

	alumno.detalle()
}
