package main

import "errors"

func intersectSlices(slice1 []int, slice2 []int) ([]int, error) {
	if len(slice1) == 0 || len(slice2) == 0 {
		return nil, errors.New("um dos slices está vazio")
	}

	intersection := make([]int, 0)
	elements := make(map[int]bool)

	// Verifica os elementos do primeiro slice
	for _, num := range slice1 {
		elements[num] = true
	}

	// Verifica os elementos do segundo slice que também estão presentes no primeiro
	for _, num := range slice2 {
		if elements[num] {
			intersection = append(intersection, num)
		}
	}

	return intersection, nil
}