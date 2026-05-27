package main

import (
	"fmt"
	"strings" // https://golang.org/pkg/strings
)

func StringFunctions() {

	// Contains
	fmt.Println("String contains", strings.Contains("Working with string functions", "functions"))
	fmt.Println("String does not contain", strings.Contains("Working with string functions", "phunctions"))

	// replace
	fmt.Println("String replace", strings.Replace("Working with string functions", "functions", "variables", -1))

	// Title case
	fmt.Println("String contains", strings.Title("Working with string functions"))

	// trim
	fmt.Println("String contains", strings.Trim("___Working with string functions___", "_"))
}
