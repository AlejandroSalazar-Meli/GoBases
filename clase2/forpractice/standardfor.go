package forpractice

import "fmt"

func StandardFor() {
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println("suma del standard for es: ", sum)
}
