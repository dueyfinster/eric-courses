package main

import "fmt"

func add(vals ...int) {
	total := 0
	for _, val := range vals {
		total += val
	}
	// return total
	fmt.Println("Variadic", vals, total)

}

// Variadic function is one that accepts many args
func VariadicFunctions() {
	vals := []int{10, 20}
	add(vals...)
	// Show we can pass more args
	vals = []int{20, 30, 40, 50}
	add(vals...)

	// just pass values instead of slices
	add(1, 2, 3, 4)
	add(1, 2, 3, 4, 5, 6)
}
