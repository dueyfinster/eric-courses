package main

import (
	"fmt"
	"math"

	"github.com/rccsdevops/go/repo"
)

func hello() {
	fmt.Println(math.Pi)
	fmt.Println(repo.SayHi())

	msg := "a string"
	fmt.Println(msg)
}
