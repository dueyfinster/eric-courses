package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
)

// Errors are values returned from functions like any other type
/* type error interface {
 Error() string
} */

// Custom errors can implement type above
// errors.New()

func HandlingErrorsInGo() {
	OpenFile()
	OpenFile2()
}

func OpenFile() {
	// func Open(name string) (file *File, err error)
	f, err := os.Open("myFile.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file successfully opened:", f.Name())
}

// return error if negative radius
func Volume(r float64) (float64, error) {
	if r < 0 {
		return 0, errors.New("Volume calculation failed; radius negative")
	}

	return ((4.0 / 3.0) * math.Pi * r * r * r), nil
}

// Defer, Panic and Recover - avoids try/catch/finally

func OpenFile2() {
	// func Open(name string) (file *File, err error)
	f, err := os.Open("myFile2.txt")
	defer f.Close()

	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("file successfully opened:", f.Name())
}
