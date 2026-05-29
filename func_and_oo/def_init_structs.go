package main

import "fmt"

// Struct is lightweight named fields, weak OO
// Box is an exportable object
type Box1 struct {
	D float64
	W float64
	H float64
}

func DefineAndInitStructs() {
	b := Box1{5, 4, 4}
	fmt.Println(b)

	// using attribute names
	c := Box1{H: 5, W: 4, D: 4}
	fmt.Println(c)

	// changing values via attributes
	d := Box1{H: 5, W: 4, D: 4}
	d.D = 6
	fmt.Println(d)

	// access attributes via pointers with explicit dererference
	ptr := &d
	(*ptr).D = 7
	fmt.Println(d)

	// access attributes via pointers without explicit dererference
	ptr2 := &d
	ptr2.D = 9
	fmt.Println(d)
}
