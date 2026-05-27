package main

import "fmt"

func DataTypes() {
	// signed integers
	var i int = -42 // cpu arch dependent
	var i8 int8 = -8
	var i16 int16 = -16
	var i32 int32 = -32 // same as a rune
	var i64 int64 = -64
	fmt.Println("int:", i)
	fmt.Println("int8:", i8)
	fmt.Println("int16:", i16)
	fmt.Println("int32:", i32)
	fmt.Println("int64:", i64)

	// unsigned integers
	var u uint = 42  // cpu arch dependent
	var u8 uint8 = 8 // same as a byte
	var u16 uint16 = 16
	var u32 uint32 = 32
	var u64 uint64 = 64
	fmt.Println("uint:", u)
	fmt.Println("uint8:", u8)
	fmt.Println("uint16:", u16)
	fmt.Println("uint32:", u32)
	fmt.Println("uint64:", u64)

	// aliases
	var byt byte = 'A' // alias for uint8
	var rn rune = '世'  // alias for int32
	fmt.Println("byte:", byt)
	fmt.Println("rune:", rn)

	var uiptr uintptr = 0xc82000c290 // cpu arch dependent
	fmt.Println("uintptr", uiptr)

	// floating point
	var f32 float32 = 3.14
	var f64 float64 = 3.1415926535
	fmt.Println("float32:", f32)
	fmt.Println("float64:", f64)

	// complex numbers - real and imaginary components
	var c64 complex64 = 1 + 2i
	var c128 complex128 = 3 + 4i
	fmt.Println("complex64:", c64)
	fmt.Println("complex128:", c128)

	// Operators
	var x int = 5
	var y int = 2
	fmt.Println("Addition", x+y)
	fmt.Println("Subtraction", x-y)
	fmt.Println("Multiplication", x*y)
	fmt.Println("Division", x/y)
	fmt.Println("Remainder", x%y)

	// bool
	var b bool = true
	fmt.Println("bool:", b)

	// Boolean Operators
	fmt.Println("bool and:", b && true)
	fmt.Println("bool or:", b || true)
	fmt.Println("bool not:", !b)

	// string - immutable
	var s string = "hello" // sequence of bytes
	fmt.Println("string:", s)
	fmt.Println("string len():", len(s)) // counts bytes - never negative
	// will not work as strings are immutable
	// s[5] = x
}
