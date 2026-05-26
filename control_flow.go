package main

import "fmt"

func ControlFlow() {
	if score := 99; score < 100 {
		fmt.Println("oops!")
	} else if score >= 100 && score < 1000 {
		fmt.Println("good!")
	} else {
		fmt.Println("wow!")
	}
}
