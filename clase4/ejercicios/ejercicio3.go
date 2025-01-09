package ejercicios

func Ejercicio3(categoria string, minutosTrabajados int) float64 {

	switch categoria {
	case "C":
		return (calcularCantidadHoras(minutosTrabajados) * 1000)
	case "B":
		return ((calcularCantidadHoras(minutosTrabajados) * 1500) * 1.2)
	case "A":
		return ((calcularCantidadHoras(minutosTrabajados) * 3000) * 1.5)
	default:
		return 0.0
	}

}

func calcularCantidadHoras(cantidadMinutos int) float64 {
	return float64(cantidadMinutos) / 60
}
