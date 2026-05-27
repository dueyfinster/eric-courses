package main

import "fmt"

// Stores memory address of a value
func Pointers() {
	//var <name> *T
	var ptr *int
	fmt.Println(ptr) // <nil>

	val := 123
	var ptr2 *int = &val
	fmt.Println(ptr2, *ptr2) // now has an address and value

	// var <name> *T = new(T)
	var ptr3 *int = new(int)
	fmt.Println(ptr3, *ptr3) // defaults to 0
	*ptr3 = 123              // changes the value
	fmt.Println(ptr3, *ptr3) // now has a value

	val2 := 123
	ptr4 := &val2            // points to same address
	fmt.Println(ptr4, *ptr4) // now has a value

	var val3 = 123
	var ptr5 = &val3
	fmt.Println(ptr5, *ptr5) // now has a value
}
