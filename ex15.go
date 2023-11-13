package main

import "fmt"

func estaPresente(x int, y []int) (bool, error) {

	estaPresenteSlice := false

	if len(y) == 0 {
		return estaPresenteSlice, fmt.Errorf("seu slice está vazio")
	}
	for _, ranY := range y {
		if ranY == x {
			estaPresenteSlice = true
			break
		} else {
			estaPresenteSlice = false
		}
	}

	return estaPresenteSlice, nil
}

func main() {
	x := 5
	y := []int{34, 3, 4, 5, 6, 7}
	z := []int{432423, 4, 8, 6, 4}
	vazio := []int{}

	fmt.Printf("Seu valor %d está presente em %d? ", x, y)
	fmt.Println(estaPresente(x, y))
	fmt.Printf("Seu valor %d está presente em %d? ", x, z)
	fmt.Println(estaPresente(x, z))
	fmt.Printf("Seu valor %d está presente em %d? ", x, vazio)
	fmt.Println(estaPresente(x, vazio))
}