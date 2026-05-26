package main

import "fmt"

func FlowControl() {

	// for loop
	sum := 0
	for i := 0; i < 5; i++ {
		sum += 1
	}
	fmt.Println("For loop", sum)

	// if statement
	temp := -10
	if temp < 0 {
		fmt.Println("Below freezing!")
	} else if temp == 0 {
		fmt.Println("At freezing point!")
	} else {
		fmt.Println("Above freezing!")
	}

	// if with scoped assignment
	if temp := 10; temp < 0 {
		fmt.Println("Below freezing!")
	} else if temp == 0 {
		fmt.Println("At freezing point!")
	} else {
		fmt.Println("Above freezing!")
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
	temp = 20
	switch {
	case temp < 0:
		fmt.Println("Below freezing!")
	case temp == 0:
		fmt.Println("At freezing point!")
	default:
		fmt.Println("Above freezing!")
	}

	// defer statement
	// arguments evaluated straight away
	// function called after parent
	defer fmt.Println("second")
	fmt.Println("first")

}
