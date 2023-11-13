package main

import "errors"

func aplicarFuncaoEmSlice(sliceInteiros []int, funcao func(int) int) ([]int, error) {
	if len(sliceInteiros) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	resultado := make([]int, len(sliceInteiros))
	for i, valor := range sliceInteiros {
		resultado[i] = funcao(valor)
	}

	return resultado, nil
}