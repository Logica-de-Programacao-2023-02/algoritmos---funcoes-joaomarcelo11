package main

import (
	"errors"
	"strings"
	"unicode"
)

func obterStringsIniciaisMaiusculas(sliceStrings []string) (string, error) {
	if len(sliceStrings) == 0 {
		return "", errors.New("O slice está vazio")
	}

	var result strings.Builder
	for _, str := range sliceStrings {
		if len(str) > 0 && unicode.IsUpper(rune(str[0])) {
			result.WriteString(str)
			result.WriteString(" ")
		}
	}

	return strings.TrimSpace(result.String()), nil
}