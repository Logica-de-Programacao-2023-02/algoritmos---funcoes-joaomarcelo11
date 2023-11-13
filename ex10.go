package main

import (
	"errors"
	"fmt"
	"sort"
)

func ordenarSlice(sliceInteiros []int) ([]int, error) {
	if len(sliceInteiros) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	sliceOrdenado := make([]int, len(sliceInteiros))
	copy(sliceOrdenado, sliceInteiros)
	sort.Ints(sliceOrdenado)

	return sliceOrdenado, nil
}

func main() {
	slice := []int{9, 2, 5, 1, 7}

	sliceOrdenado, err := ordenarSlice(slice)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("Slice ordenado:", sliceOrdenado)
}