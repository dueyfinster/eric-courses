package main

import "fmt"

func Variables() {
	var msg string = "a string"
	fmt.Println("declared string", msg)

	var msg2 string
	fmt.Println("empty string", msg2)

	var msg3 string
	msg3 = "a string"
	fmt.Println("declare after init string", msg3)

	// type is inferred by compiler
	msg4 := "a string"
	fmt.Println("short declare string", msg4)

	var x int
	var truth bool
	var msg5 string
	var pointer *string
	fmt.Println("integer: ", x)
	fmt.Println("boolean: ", truth)
	fmt.Println("string: ", msg5)
	fmt.Println("pointer: ", pointer)

}
