package main

import (
	"fmt"
	"strings"
)
func calcularMedia(nums []int) float64 {
	soma := 0

	for _, num := range nums {
		soma += num
	}

	return float64(soma) / float64(len(nums))
}

func contarVogais(s string) int {
	vogais := []string{"a", "e", "i", "o", "u"}
	count := 0

	for _, char := range s {
		if contains(vogais, strings.ToLower(string(char))) {
			count++
		}
	}

	return count
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func concatenarStrings(strs []string) string {
	return strings.Join(strs, "")
}

func encontrarSegundoMaior(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	maior := nums[0]
	segundoMaior := nums[1]

	for i := 1; i < len(nums); i++ {
		if nums[i] > maior {
			segundoMaior = maior
			maior = nums[i]
		} else if nums[i] > segundoMaior {
			segundoMaior = nums[i]
		}
	}

	return segundoMaior
}

func encontrarPosicao(nums []int, valor int) int {
	for i, num := range nums {
		if num == valor {
			return i
		}
	}

	return -1
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	strs := []string{"Hello", " ", "World", "!"}

	media := calcularMedia(nums)
	fmt.Println("Média:", media)

	vogais := contarVogais("Hello World!")
	fmt.Println("Vogais:", vogais)

	concatenado := concatenarStrings(strs)
	fmt.Println("Concatenado:", concatenado)

	segundoMaior := encontrarSegundoMaior(nums)
	fmt.Println("Segundo maior:", segundoMaior)

	posicao := encontrarPosicao(nums, 3)
	fmt.Println("Posição:", posicao)
}