package main

import "fmt"

// globals - shouldn't use
// var x int = 20
// var y int = 10

// using individdual variables
// func add(x, y, z int) int {
// 	return x + y + z
// }

func add(x []int) int {
	total := 0

	for _, val := range x {
		total += val
	}
	return total
}

func mul(x, y *int) int {
	return *x * *y
}

func PassingParamsToFunc() {
	// using individual values
	// x := 20
	// y := 10
	// z := 5
	//
	// fmt.Println(add(x, y, z))

	s := []int{10, 20, 5}
	fmt.Println(add(s))

	x := 5
	y := 2
	fmt.Println(mul(&x, &y))
}
