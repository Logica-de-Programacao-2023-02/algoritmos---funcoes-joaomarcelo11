package main

import (
	"fmt"
	"strings"
)

func contarVogais(s string) int {
	vogais := "aeiouAEIOU"
	count := 0
	for _, char := range s {
		if strings.ContainsAny(string(char), vogais) {
			count++
		}
	}
	return count
}
func concatenarStrings(strs []string) string {
	return strings.Join(strs, "")
}

func segundoMaiorValor(nums []int) int {
	if len(nums) < 2 {
		return -1
	}
	maior, segundoMaior := nums[0], nums[1]
	if segundoMaior > maior {
		maior, segundoMaior = segundoMaior, maior
	}
	for i := 2; i < len(nums); i++ {
		if nums[i] > maior {
			maior, segundoMaior = nums[i], maior
		} else if nums[i] > segundoMaior {
			segundoMaior = nums[i]
		}
	}
	return segundoMaior
}

func encontrarIndiceValor(nums []int, valor int) int {
	for i, num := range nums {
		if num == valor {
			return i
		}
	}
	return -1
}

func main() {
	s1 := "hello world"
	fmt.Printf("A string \"%s\" possui %d vogais.\n", s1, contarVogais(s1))

	strs := []string{"foo", "bar", "baz"}
	fmt.Printf("A concatenação das strings \"%v\" é \"%s\".\n", strs, concatenarStrings(strs))

	nums := []int{1, 2, 3, 4, 5}
	fmt.Printf("O segundo maior valor em %v é %d.\n", nums, segundoMaiorValor(nums))

	nums2 := []int{10, 20, 30, 40, 50}
	valor := 30
	fmt.Printf("O índice do primeiro elemento igual a %d em %v é %d.\n", valor, nums2, encontrarIndiceValor(nums2, valor))
}