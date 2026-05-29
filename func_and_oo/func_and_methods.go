package main

import "fmt"

func FuncAndMethods() {
	y := []float64{1.2, 3.4, 5.6}
	fmt.Println("Average:", avg(y))
	c := Cube{

		w: 10,
		h: 20,
		d: 30,
	}
	// note .volume() method attched to Cube type
	fmt.Println("Cube", c.volume())
}

func avg(x []float64) float64 {
	total := 0.0
	for _, val := range x {
		total += val
	}
	return total / float64(len(x))
}

// define a volume method for a cube
// note Cube reciever at the start
func (c *Cube) volume() float64 {
	return c.d * c.w * c.h
}

type Cube struct {
	w float64
	h float64
	d float64
}
