package main

import (
	"fmt"

	"github.com/AlejandroSalazar-Meli/clase5/ejercicios"
)

type Product struct {
	ID          int
	Name        string
	Price       int
	Description string
	Category    string
}

func main() {
	var audifonosSony ejercicios.Product = ejercicios.Product{
		Name:        "Audifonos Sony XMH4",
		ID:          004,
		Price:       3000,
		Description: "Audifonos sony de buena calidad",
		Category:    "Audio",
	}

	audifonosSony.Save()
	audifonosSony.GetAll()
	productoEscogido, err := audifonosSony.GetById(004)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El producto escogido es: %v", productoEscogido)
	}
}
