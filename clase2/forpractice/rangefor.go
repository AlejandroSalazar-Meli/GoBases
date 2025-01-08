package forpractice

import "fmt"

func RangeFor() {
	paises := []string{"Colombia", "Argentina", "Mexico", "Uruguay", "Chile"}
	for i, pais := range paises {
		fmt.Println(i, pais)
	}
}
