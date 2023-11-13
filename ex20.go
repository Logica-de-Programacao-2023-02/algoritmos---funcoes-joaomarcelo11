package main

import (
	"fmt"
)

func apenasMaioresQueCinco(x []string) ([]string, error) {

	maioresQueCinco := []string{}

	if len(x) == 0 {
		return x, fmt.Errorf("seu slice está vazio")
	}

	for _, ranX := range x {

		if len(ranX) > 5 {
			maioresQueCinco = append(maioresQueCinco, ranX)
		}

	}

	return maioresQueCinco, nil

}

func main() {

	x := []string{"Isabelle", "Gabriel", "Tutti", "Bolt", "Baby", "Frajolinha"}
	y := []string{}

	fmt.Printf("Seus strings maiores que 5 em %s são: ", x)
	fmt.Println(apenasMaioresQueCinco(x))
	fmt.Printf("Seus strings maiores que 5 em %s são: ", y)
	fmt.Println(apenasMaioresQueCinco(y))

}