package main

import (
	"errors"
	"fmt"
	"strconv"
)

func somarDigitos(numero int) (int, error) {
	if numero < 0 {
		return 0, errors.New("O número é negativo")
	}

	strNumero := strconv.Itoa(numero)
	soma := 0
	for _, char := range strNumero {
		digito, err := strconv.Atoi(string(char))
		if err != nil {
			return 0, err
		}
		soma += digito
	}

	return soma, nil
}

func main() {
	numero := 12345

	somaDigitos, err := somarDigitos(numero)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("Soma dos dígitos:", somaDigitos)
}