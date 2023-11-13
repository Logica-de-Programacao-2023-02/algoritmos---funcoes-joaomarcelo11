package main

import (
	"fmt"
	"strings"
)

func concatenarComVirgulas(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", fmt.Errorf("O slice está vazio")
	}

	resultado := strings.Join(slice, ", ")
	return resultado, nil
}