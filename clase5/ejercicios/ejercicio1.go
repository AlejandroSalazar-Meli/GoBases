package ejercicios

import (
	"errors"
	"fmt"
)

type Product struct {
	ID          int
	Name        string
	Price       int
	Description string
	Category    string
}

var Products = []Product{
	{
		ID:          001,
		Name:        "MacBook pro",
		Price:       1000,
		Description: "Computador portatil de apple",
		Category:    "Computadores",
	},
	{
		ID:          002,
		Name:        "Ipad mini",
		Price:       500,
		Description: "Tablet de apple",
		Category:    "Tablets",
	},
	{
		ID:          003,
		Name:        "Iphone 13 pro max",
		Price:       650,
		Description: "Celular de apple",
		Category:    "Celulares",
	},
}

func (product Product) Save() {
	Products = append(Products, product)
}

func (product Product) GetAll() {
	for i, prod := range Products {
		fmt.Println(prod)
		fmt.Printf("Producto %d: \nID: %d\nNombre: %s\nPrecio: %d\nDescripción: %s\nCategoría: %s\n", i, prod.ID, prod.Name, prod.Price, prod.Description, prod.Category)
	}
}

func (product Product) GetById(identificador int) (Product, error) {
	for _, prod := range Products {
		if prod.ID == identificador {
			return prod, nil
		}
	}
	return Product{}, errors.New("no se encontró el producto con ese ID")
}
