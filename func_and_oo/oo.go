package main

import "fmt"

// Go uses structs instead
// structs have named fields and nethods optionally
func OO() {
	// Objects
	c := Cube1{depth: 4, width: 4, height: 4}
	fmt.Println("Cube vol:", c.volume())

	// Inheritance
	t := Tire{Part{Manufacturer: "Brocadero"}}
	// first version
	// fmt.Println("Tyres:", t.Part.Mfc())
	fmt.Println("Tyres:", t.Mfc())

	// Interfaces
	sp := Sphere1{radius: 4}
	var shapes []Shape1

	shapes = append(shapes,
		c,
		sp,
	)

	fmt.Println("Total volume:", totalVolume1(shapes...))
}

type Cube1 struct {
	width  float64
	height float64
	depth  float64
}

type Sphere1 struct {
	radius float64
}

type Shape1 interface {
	volume() float64
}

func (c Cube1) volume() float64 {
	return c.depth * c.width * c.height
}

// 4/3 * pi * r*r*r
func (sp Sphere1) volume() float64 {
	return ((4.0 / 3.0) * 3.14 * (sp.radius * sp.radius * sp.radius))
}

func totalVolume1(shapes ...Shape1) float64 {
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
