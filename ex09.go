package main

import (
	"errors"
	"fmt"
	"strings"
)

func obterPalavrasDaString(texto string) ([]string, error) {
	if texto == "" {
		return nil, errors.New("A string está vazia")
	}

	palavras := strings.Split(texto, " ")
	return palavras, nil
}

func main() {
	texto := "Olá, mundo! Este é um exemplo de string."

	palavras, err := obterPalavrasDaString(texto)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("Palavras:", palavras)
}