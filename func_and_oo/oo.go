package main

import "fmt"

// Go uses structs instead
// structs have named fields and nethods optionally
func OO() {
	// Objects
	c := Cube{depth: 4, width: 4, height: 4}
	fmt.Println("Cube vol:", c.volume())

	// Inheritance
	t := Tire{Part{Manufacturer: "Brocadero"}}
	// first version
	// fmt.Println("Tyres:", t.Part.Mfc())
	fmt.Println("Tyres:", t.Mfc())

	// Interfaces
	sp := Sphere{radius: 4}
	var shapes []Shape

	shapes = append(shapes,
		c,
		sp,
	)

	fmt.Println("Total volume:", totalVolume(shapes...))
}

type Cube struct {
	width  float64
	height float64
	depth  float64
}

type Sphere struct {
	radius float64
}

type Shape interface {
	volume() float64
}

func (c Cube) volume() float64 {
	return c.depth * c.width * c.height
}

// 4/3 * pi * r*r*r
func (sp Sphere) volume() float64 {
	return ((4.0 / 3.0) * 3.14 * (sp.radius * sp.radius * sp.radius))
}

func totalVolume(shapes ...Shape) float64 {
	var volume float64

	for _, s := range shapes {
		volume += s.volume()
	}

	return volume
}

// Inheritance more complicated but achieved using embedded types
type Part struct {
	Manufacturer string
}

func (p *Part) Mfc() string {
	return p.Manufacturer
}

// first version with a type Part
// type Tire struct {
// 	Part Part
// }

// second version with anonymous type
type Tire struct {
	Part // anonymous field
}
