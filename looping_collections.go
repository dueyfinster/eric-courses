package main

import "fmt"

func LoopingCollections() {

	// looping a slice
	sl := []int{10, 20, 30, 40}
	for i := 0; i < len(sl); i++ {
		fmt.Println("Looping slice", i, sl[i])
	}

	// same as above but rewritten to use range
	for i, value := range sl {
		fmt.Println("Looping slice with range", i, value)
	}

	prodPrice := map[string]int{
		"widget":             75,
		"turbo widget":       100,
		"convertible widget": 150,
	}

	for key, value := range prodPrice {
		fmt.Println("Looping map with range", key, value)
	}

	// just keys
	for key := range prodPrice {
		fmt.Println("Looping map with keys only", key)
	}

	// just values
	for _, value := range prodPrice {
		fmt.Println("Looping map with values only", value)
	}
}
