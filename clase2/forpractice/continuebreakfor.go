package forpractice

import "fmt"

func ContinueBreakFor() {
	cantidadPares := 0
	for i := 0; i <= 100; i++ {
		if i%2 != 0 {
			continue
		}
		cantidadPares += 1
	}

	fmt.Println("La cantidad de pares de 0 a 100 es: ", cantidadPares)

	cantidadImparesHastaSumarMil := 0
	sumaImpares := 0
	i := 0

	for {
		if i%2 != 0 {
			cantidadImparesHastaSumarMil += 1
			sumaImpares += i
		}
		if sumaImpares >= 1000 {
			break
		}
		i++
	}

	fmt.Println("La cantidad de impares hasta sumar mil es: ", cantidadImparesHastaSumarMil, " y la suma es: ", sumaImpares)
}
