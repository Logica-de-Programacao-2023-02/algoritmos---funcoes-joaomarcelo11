package main

import "fmt"

func primosAteInt(x int) ([]int, error) {

	primoSlice := true
	slicePrimos := []int{}

	if x < 0 {
		return slicePrimos, fmt.Errorf("seu número é negativo")
	}

	for i := 2; i <= x; i++ {

		primoSlice = true

		for i2 := 2; i2 < i; i2++ {

			if i%i2 == 0 {

				primoSlice = false

			}

		}

		if primoSlice {
			slicePrimos = append(slicePrimos, i)
		}
	}

	return slicePrimos, nil

}

func main() {

	x := 79
	y := -1

	fmt.Printf("Todos os números primos até %d são: ", x)
	fmt.Println(primosAteInt(x))
	fmt.Printf("Todos os números primos até %d são: ", y)
	fmt.Print(primosAteInt(y))

}