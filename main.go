package main

import (
	"fmt"
	"generic/interface_example"
	mapExample "generic/map_example"

	"golang.org/x/exp/constraints"
)

func mapValues[T constraints.Ordered](values []T, fn func(T) T) []T {
	newValues := make([]T, 0)
	for _, data := range values {
		result := fn(data)
		newValues = append(newValues, result)
	}
	return newValues
}
func main() {
	resultWithInt := mapValues([]int{1, 2, 3}, func(n int) int {
		return n * 2
	})
	fmt.Println(resultWithInt)

	resultWithFloat := mapValues([]float64{1.1, 2.2, 3.3}, func(n float64) float64 {
		return n * 2
	})
	fmt.Println(resultWithFloat)

	//! Example with interface
	interface_example.ExampleWithInterface("nahkar")

	//! Example with map
	mapExample.ExampleWithMap()
}
