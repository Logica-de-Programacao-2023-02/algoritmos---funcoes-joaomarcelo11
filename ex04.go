package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func secondLargest(numbers []int) (int, error) {
	if len(numbers) < 2 {
		return 0, errors.New("Slice must have at least 2 elements")
	}

	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	return numbers[1], nil
}

func findIndex(numbers []int, value int) int {
	for i, n := range numbers {
		if n == value {
			return i
		}
	}

	return -1
}

func joinStrings(strings []string) (string, error) {
	if len(strings) == 0 {
		return "", errors.New("Slice must have at least 1 element")
	}

	return strings.Join(strings, ","), nil
}

func main() {
	numbers := []int{10, 5, 8, 12, 3}
	fmt.Println("Slice de números:", numbers)
	second, err := secondLargest(numbers)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Segundo maior número:", second)
	}

	index := findIndex(numbers, 8)
	fmt.Println("Índice do valor 8:", index)

	strings := []string{"hello", "world", "how", "are", "you"}
	fmt.Println("Slice de strings:", strings)
	joined, err := joinStrings(strings)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Strings concatenadas:", joined)
	}
}