package main

import "fmt"

func ControlFlow() {
	// declare variable in if scope
	if score := 99; score < 100 {
		fmt.Println("oops!")
	} else if score >= 100 && score < 1000 {
		fmt.Println("good!")
	} else {
		fmt.Println("wow!")
	}

	// Case statements
	// automatically includes "break"
	workday := 3
	switch workday {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("Take the weekend off!")
	}

	// switch with conditions
	// compare to if statement above
	score := 100
	switch {
	case score < 100:
		fmt.Println("oops!")
	case score >= 100:
		fmt.Println("good!")
	default:
		fmt.Println("wow!")
	}
}
