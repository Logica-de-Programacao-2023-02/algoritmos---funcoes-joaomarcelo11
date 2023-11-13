package main

import (
	"fmt"
)

func aplicarFuncEmIntSlice(x []int, y func(int) (int, error)) (int, error) {

	funcRanX := 0
	somaFuncRanX := 0

	if len(x) == 0 {
		return 0, fmt.Errorf("seu slice está vazio")
	}

	for _, ranX := range x {
		funcRanX, _ = y(ranX)
		somaFuncRanX += funcRanX
	}

	return funcRanX, nil

}