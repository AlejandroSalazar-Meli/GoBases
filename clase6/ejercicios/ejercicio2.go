package ejercicios

import "fmt"

type Product interface {
	Price() float64
}

type Small struct {
	Precio float64
}

func (s Small) Price() float64 {
	return s.Precio
}

type Medium struct {
	Precio float64
}

func (m Medium) Price() float64 {
	return m.Precio + (m.Precio * 0.03)
}

type Large struct {
	Precio float64
}

func (l Large) Price() float64 {
	return l.Precio + (l.Precio * 0.06) + 2500
}

func productFactory(tipoProducto string, precio float64) Product {
	switch tipoProducto {
	case "Small":
		return Small{Precio: precio}
	case "Medium":
		return Medium{Precio: precio}
	case "Large":
		return Large{Precio: precio}
	default:
		return nil
	}
}

func EjercicioDos() {
	productoPequeno := productFactory("Small", 2000.0)
	productoMediano := productFactory("Medium", 2000.0)
	productoGrande := productFactory("Large", 2000.0)

	fmt.Printf("El precio total del producto pequeño es: %2f\nEl precio total del producto mediano es: %2f\nEl precio total del producto grande es: %2f", productoPequeno.Price(), productoMediano.Price(), productoGrande.Price())

}
