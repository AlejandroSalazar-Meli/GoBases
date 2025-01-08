package forpractice

import "fmt"

func WhileFor() {
	sum := 1
	iterations := 0
	for sum < 100 {
		sum += sum
		iterations += 1
	}
	fmt.Println("La cantidad de iteraciones para que sum fuera mayor a 100 fue: ", iterations)
}
