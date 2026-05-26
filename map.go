package main

import "fmt"

func Maps() {
	var prodPrice map[string]int
	fmt.Println("prodPrice empty", prodPrice) // empty map

	//declare and initialise with make
	tempPrice := make(map[string]int)
	tempPrice["convertible widget"] = 150
	prodPrice = tempPrice
	prodPrice["widget"] = 100
	fmt.Println("prodPrice", prodPrice)

	//declare and initialise map literal
	empPrice := map[string]int{
		"widget": 75,
	}
	empPrice["turbo widget"] = 100
	fmt.Println("empPrice", empPrice)

	// check for element that exists in the map
	el, ok := empPrice["widget"]
	fmt.Println("existing el", el, "ok", ok)

	// check for element that does not exist in the map
	el1, ok1 := empPrice["wydget"]
	// will return default value of int because map uses it
	fmt.Println("not existing el1", el1, "ok1", ok1)

	fmt.Println("len() of empPrice", len(empPrice))
	delete(empPrice, "widget")
	// check for element that has been deleted in the map
	el2, ok2 := empPrice["widget"]
	fmt.Println("existing el2", el2, "ok2", ok2)
	fmt.Println("len() of empPrice", len(empPrice))
}
