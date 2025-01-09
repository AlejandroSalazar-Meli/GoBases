package ejercicios

import "fmt"

const (
	perro     = "perro"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func Ejercicio5(cantidadAnimal int) (cantidadAlimentoPerro float64, cantidadAlimentoGato float64, cantidadAlimentoHamster float64, cantidadAlimentoTarantula float64) {
	funcPerro, err1 := Animal(perro)
	funcGato, err2 := Animal(gato)
	funcHamster, err3 := Animal(hamster)
	funcTarantula, err4 := Animal(tarantula)
	funcError, err5 := Animal("animal")

	cantidadPerros := cantidadAnimal
	cantidadGatos := cantidadAnimal
	cantidadHmasters := cantidadAnimal
	cantidadTarantulas := cantidadAnimal

	if err1 != "" {
		fmt.Println(err1)
	} else {
		fmt.Printf("La cantidad de alimento para %d perros es: %.1f kilogramos\n", cantidadPerros, funcPerro(cantidadPerros))
		cantidadAlimentoPerro = funcPerro(cantidadPerros)
	}

	if err2 != "" {
		fmt.Println(err2)
	} else {
		fmt.Printf("La cantidad de alimento para %d gatos es: %.1f kilogramos\n", cantidadGatos, funcGato(cantidadGatos))
		cantidadAlimentoGato = funcGato(cantidadGatos)
	}

	if err3 != "" {
		fmt.Println(err3)
	} else {
		fmt.Printf("La cantidad de alimento para %d hamsters es: %.1f kilogramos\n", cantidadHmasters, funcHamster(cantidadHmasters))
		cantidadAlimentoHamster = funcHamster(cantidadHmasters)
	}

	if err4 != "" {
		fmt.Println(err4)
	} else {
		fmt.Printf("La cantidad de alimento para %d tarantulas es: %.1f kilogramos\n", cantidadTarantulas, funcTarantula(cantidadTarantulas))
		cantidadAlimentoTarantula = funcTarantula(cantidadTarantulas)
	}

	if err5 != "" {
		fmt.Println(err5)
	} else {
		fmt.Printf("La cantidad de alimento para %d perros es: %.1f kilogramos\n", cantidadPerros, funcError(cantidadPerros))
	}
	return
}

func calcularAlimentoPerro(cantidadAnimal int) float64 {
	return float64(cantidadAnimal) * 10
}

func calcularAlimentoGato(cantidadAnimal int) float64 {
	return float64(cantidadAnimal) * 5
}

func calcularAlimentoHamster(cantidadAnimal int) float64 {
	return float64(cantidadAnimal) * 0.250
}

func calcularAlimentoTarantula(cantidadAnimal int) float64 {
	return float64(cantidadAnimal) * 0.150
}

func Animal(tipoAnimal string) (funcionCalcular func(cantidadAnimal int) float64, err string) {
	switch tipoAnimal {
	case perro:
		funcionCalcular = calcularAlimentoPerro
	case gato:
		funcionCalcular = calcularAlimentoGato
	case hamster:
		funcionCalcular = calcularAlimentoHamster
	case tarantula:
		funcionCalcular = calcularAlimentoTarantula
	default:
		err = "Animal " + tipoAnimal + " inválido"
	}
	return
}