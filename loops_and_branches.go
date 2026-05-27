package main

import "fmt"

func LoopsAndBranches() {
	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println("For loop", i)
	}

	// for loop init i outside - simplified condition
	i := 0
	for i < 5 {
		fmt.Println("For loop2", i)
		i++
	}

	// break statement
	for i := 0; i < 5; i++ {
		if i == 2 {
			break
		}
		fmt.Println("For loop3", i)
	}

	// continue statement
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue // won't print 2 - skips iteration
		}
		fmt.Println("For loop3", i)
	}

	// goto statement
	i = 0
outerlabel:
	for i < 5 {
		if i == 2 {
			i++
			goto outerlabel // skips 2
		}
		fmt.Println("For loop4", i)
		i++
	}

	// infinite loop
	var x int
	for {
		if x == 3 {
			break
		}
		x++
		fmt.Println("For loop5", x)
	}
}
