package main

import (
	"fmt"
)

func SumInts(m map[string]int64) int64 {
	var total int64
	for _, val := range m {
		total += val
	}
	return total
}

func SumFloats(m map[string]float64) float64 {
	var total float64
	for _, val := range m {
		total += val
	}
	return total
}

// "comparable" is an interface for all types on which != && == can be used
// This is needed because go requires map keys to be comparable values
func SumIntsOrFloat[K comparable, V int64 | float64](m map[K]V) V {
	var total V
	for _, m := range m {
		total += m;
	}
	return total
}

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats),
	)
	
	fmt.Printf("Generic Sums: %v and %v\n",
		// The type arguments can be omitted and inferred by the compiler
		SumIntsOrFloat[string, int64](ints),
		SumIntsOrFloat[string, float64](floats),
	)
}
