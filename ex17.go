package main

import (
	"fmt"
	"strings"
)

func ordemAlfabetica(x []string) (string, error) {

	if len(x) == 0 {
		return "", fmt.Errorf("slice vazio")
	}

	newString := ""

	for i := 0; i < len(x); i++ {

		capsX := strings.ToUpper(x[i])

		for j := 0; j < len(x); j++ {

			capsJ := strings.ToUpper(x[j])

			if capsJ > capsX {

				x[i], x[j] = x[j], x[i]

			}
		}
	}

	for _, ranX := range x {

		newString += ranX

	}

	return newString, nil

}

func main() {
	x := []string{"uepaaaa", "Rapaz", "Isabelle", "tangente", "hahahahaa"}
	vazio := []string{}

	fmt.Print("Sua string em ordem é alfabética: ")
	fmt.Println(ordemAlfabetica(x))
	fmt.Print("Sua string em ordem é alfabética: ")
	fmt.Println(ordemAlfabetica(vazio))

}