package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// Working With Errors Package
// https://pkg.go.dev/errors

// See where Error interface is defined:
// https://pkg.go.dev/builtin#error
func WorkingPkgErrors() {

	if _, err := os.Open("myFile3.txt"); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Println("File does not exist wp error")
		} else {
			log.Println(err)
		}
		return
	}
	fmt.Println("File successfully opened")
}
