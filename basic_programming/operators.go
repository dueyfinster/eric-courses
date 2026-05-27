package main

import (
	"fmt"
	"math"
)

func Operators() {
	fmt.Println("Hello from go!")
	fmt.Println(math.Pi)
	var i int = 10
	const c string = "Yes"
	const b bool = false

	fmt.Println(i, c, b)
	Do_math()
	Do_math_rel()
	Do_math_bit()
	Assign_op()
	String_ptr()
}

func Do_math() {
	x := 5
	y := 2
	fmt.Println(x-y, x*y, x+y, x/y, x%y)
}

func Do_math_rel() {
	x := 5
	y := 2
	fmt.Println(x == y, x != y, x < y, x > y, x >= 5)
}

func Do_math_log() {
	x := 5
	y := 2
	fmt.Println(x > y && x >= 5)
	fmt.Println(x > y || x < 5)
	fmt.Println(!true, !false)
}

func Do_math_bit() {
	// x := 5
	// y := 2
	fmt.Println("Bitshift 2 right: ", 2<<1) // 0010 to 0100
	fmt.Println("Bitshift 2 left: ", 2>>1)  // 0010 to 0001
	fmt.Println("2 and 1: ", 2&1)           // 0010 & 0001 to 0000
	fmt.Println("2 and not 1: ", 2&^1)      // 0010 &^ 0001 to 0010
	fmt.Println("2 or 1: ", 2|1)            // 0010 | 0001 to 0011
}

func Assign_op() {
	x := 5
	y := 2
	x += y
	fmt.Println("x+=y", x)
	x -= y
	fmt.Println("x-=y", x)
	x *= y
	fmt.Println("x*=y", x)
	x /= y
	fmt.Println("x/=y", x)
}

func String_ptr() {
	x := "hello"
	y := &x
	fmt.Println(y)
	fmt.Println(*y)
	fmt.Println(x)
	*y = "goodbye"
	fmt.Println(x)
}
