package main

import "fmt"

// Go always creates copies of arrays when passing as args

func Arrays() {
	normal_arrays()
	two_d_arrays()
}

func normal_arrays() {
	var arr [4]int
	fmt.Println("arr values", arr)

	var arr2 = [4]int{0, 1, 2, 3}
	fmt.Println("arr2 values", arr2)

	var arr3 = [4]int{0, 1, 2, 3}
	arr3[1] = 10
	fmt.Println("arr3[1] value", arr3[1])
	fmt.Println("arr3 length", len(arr3))
	fmt.Println("arr3 values", arr3)
}

func two_d_arrays() {
	var twoDimArr [3][2]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			twoDimArr[i][j] = i + j
		}
	}
	fmt.Println("two dimensional array values", twoDimArr)
}
