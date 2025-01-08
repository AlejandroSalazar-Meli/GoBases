package ifelse

import "fmt"

func Ifelse() {
	edadAlejandro := 36
	edadYuliana := 34

	if edadYuliana > edadAlejandro {
		fmt.Println("Yuliana es mayor que Alejandro")
	} else {
		fmt.Println("Alejandro es mayor que Yuliana")
	}
}
