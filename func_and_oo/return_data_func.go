package main

import "fmt"

// no return version
// func add(x, y int) {
// 	fmt.Println(x + y)
// }

// just using return type
// func add(x, y int) int {
// 	return x + y
// }

func add4(x, y int) (a, b int, c bool) {
	a = x + y
	b = x - y
	c = x > y

	// default return will use methood signature
	return
	// customising return - overwriting c in this case
	// return a, b, false
}

func ReturnDataFunctions() {
	x := 20
	y := 10
	fmt.Println(add4(x, y))
}
