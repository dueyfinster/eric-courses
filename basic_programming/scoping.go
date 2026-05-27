package main

import "fmt"

var workday = 5

func Scoping() {
	i := 10
	for i := 0; i < 5; i++ {
		fmt.Println("Value inside for of i", i)
	}
	fmt.Println("Value outside of for loop of i", i)

	for i < 15 {
		fmt.Println("value of i init inside for loop", i)
		i++
	}
	fmt.Println("Value outside of for loop of i", i)

	// also if statements subject to scoping
	j := 1
	if j < 5 {
		fmt.Println("value of i inside if statement", i)
	}
	fmt.Println("Value outside of for loop of i", i)
	workdays()
}

func workdays() {
	//local scope
	// workday := 3
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
		fmt.Println("Take the weekned off!")
	}
}
