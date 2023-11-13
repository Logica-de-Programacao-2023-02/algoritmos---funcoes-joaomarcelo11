package main

import (
	"errors"
	"fmt"
)

func obterNumerosPares(sliceInteiros []int) ([]int, error) {
	if len(sliceInteiros) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	numerosPares := []int{}
	for _, valor := range sliceInteiros {
		if valor%2 == 0 {
			numerosPares = append(numerosPares, valor)
		}
	}

	return numerosPares, nil
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	numerosPares, err := obterNumerosPares(slice)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("Números pares:", numerosPares)
}