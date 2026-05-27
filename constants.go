package main

import (
	"fmt"
	"math"
)

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func Constants() {
	// constant from stdlib
	fmt.Println(math.Pi)

	const pi = 3.14
	// no possible to redefine const
	// pi = 3.14159
	fmt.Println(pi)

	// implicitly typed constant
	const myConst = 1
	// now float64 and value 2.1
	fmt.Printf("type %T, value: %v\n", myConst+1.1, myConst+1.1)

	const myConst1 int = 1
	// define float64 explicitly
	fmt.Printf("type %T, value: %v\n", float64(myConst1)+1.1, float64(myConst1)+1.1)

	// Will not work as func eval at runtime and constants need to be known at compile time
	// const myConst2 int = setMe()
	// define float64 explicitly
	// fmt.Printf("type %T, value: %v", myConst2 + 1.1, myConst2 +1.1)

	// enum with iota
	fmt.Println("enum iota:", Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}

func setMe() int {
	return 1
}
