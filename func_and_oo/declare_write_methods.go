package main

import "fmt"

// Box is an exportavle object
type Box struct {
	D float64
	W float64
	H float64
}

// volume method to calculate volume of a box
func (b *Box) volume() float64 {
	return b.D * b.W * b.H
}

func DeclareAndWriteMethods() {
	b := Box{D: 5, W: 4, H: 3}
	fmt.Println(b.volume())
}
