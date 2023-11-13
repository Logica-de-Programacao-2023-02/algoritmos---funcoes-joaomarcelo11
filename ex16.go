package main

import (
	"fmt"
	"strings"
)

func ocultarVogais(x string) (string, error) {

	capX := strings.ToUpper(x)
	sliceCapX := strings.Split(capX, "")
	newX := ""

	if len(x) == 0 {
		return x, fmt.Errorf("seu string está vazio")
	}

	for _, ranSliceCapX := range sliceCapX {
		newX += ranSliceCapX
	}

	newCapX := strings.Split(newX, "")
	sliceX := strings.Split(x, "")

	for i := 0; i < len(newCapX); i++ {

		if newCapX[i] == "A" || newCapX[i] == "E" || newCapX[i] == "I" || newCapX[i] == "O" || newCapX[i] == "U" {

			sliceX[i] = "*"

		}
	}

	xVogaisOcultas := ""
	for _, ranSliceX := range sliceX {
		xVogaisOcultas += ranSliceX
	}

	return xVogaisOcultas, nil

}

func main() {

	x := "Hello World!"
	y := ""

	fmt.Print("Sua frase com as vogais ocultas é: ")
	fmt.Println(ocultarVogais(x))
	fmt.Print("Sua frase com as vogais ocultas é: ")
	fmt.Print(ocultarVogais(y))

}