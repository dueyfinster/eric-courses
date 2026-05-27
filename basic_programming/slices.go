package main

import "fmt"

// Slices are passed by reference (not a full copy)
// Slices are NOT a fixed size
// Slices are done by type they hold
func Slices() {
	var sl []int
	fmt.Println("Slice default", sl)
	fmt.Println("is Slice nil?", sl == nil)

	//make(T, len, cap)
	var sl2 = make([]int, 4)
	fmt.Println("Slice2 default", sl2)
	fmt.Println("is Slice2 nil?", sl2 == nil)

	var sl3 = []int{0, 1, 2, 3}
	fmt.Println("Slice3 default", sl3)
	fmt.Println("is Slice3 nil?", sl3 == nil)

	sl3[1] = 10
	fmt.Println("Slice3[1]", sl3[1])
	fmt.Println("Slice3 len()", len(sl3))

	//append([]T, element1, element2)
	//append([]T, []T...)
	sl4 := []int{0, 1, 2, 3}
	sl4 = append(sl4, 4, 5, 6)
	fmt.Println("length of sl4:", len(sl4))
	fmt.Println("sl4 values:", sl4)

	sl5 := []int{0, 1, 2, 3}
	slAppend := []int{4, 5, 6, 7}
	sl5 = append(sl5, slAppend...) // have to unbox
	fmt.Println("length of sl5:", len(sl5))
	fmt.Println("sl5 values:", sl5)

	//copy(dst, src []T) int
	sl6 := []int{0, 1, 2, 3}
	slCopy := make([]int, len(sl6))
	el := copy(slCopy, sl6) // have to unbox
	fmt.Println("length of slCopy:", len(slCopy))
	fmt.Println("slCopy values:", slCopy)
	fmt.Println("number of elements copied", el)

	sl7 := []int{0, 1, 2, 3, 4, 5}
	//[]T  [low, high]
	// NOTE: high is up to but not inclusive
	sl8 := sl7[2:4] // specific range
	sl9 := sl7[:5]  // defualts to 0  on lhs
	sl10 := sl7[3:] // goes to the end
	fmt.Println("Values in sl7:", sl7)
	fmt.Println("Values in sl8:", sl8)
	fmt.Println("Values in sl9:", sl9)
	fmt.Println("Values in sl10:", sl10)

}
