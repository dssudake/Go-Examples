// Program to learn types in Golang

/*
Basic types available in go are :
	bool
	Numeric Types
	 - int8, int16, int32, int64, int
	 - uint8, uint16, uint32, uint64, uint
	 - float32, float64
	 - complex64, complex128
	 - byte
	 - rune
	string
*/
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// bool - either true or false
	var a bool = true
	b := false
	fmt.Println("boolean \na:", a, " b:", b)
	c := a && b
	d := a || b
	fmt.Println("c:", c, "d:", d)

	/*
		signed integers   - int8, int16, int32, int64, int
		unsigned integers - uint8, uint16, uint32, uint64, uint

		int & uint represents 32 or 64 bit depending upon underlying platform

		byte : alias of uint8
		rune : alias of int32

		range -->  -(2^(n-1)) to (2^(n-1))-1
		e.g.: int32 --> n=32 --> -2147483648 to 2147483647
	*/
	var ai int = 40
	bi := 60
	fmt.Println("\ninteger \na:", ai, "b:", bi)

	/*
		Type of variable can be printed using %T format specifier in Printf function
		%T - type of variable
		%d - size of variable

		'unsafe' package in go has 'Sizeof' method which return size of variable in bytes
	*/
	fmt.Printf("type of a : %T, size of a : %d", ai, unsafe.Sizeof(ai))
	fmt.Printf("\ntype of b : %T, size of b : %d", bi, unsafe.Sizeof(bi))

	/*
		floating point numbers - float32, float64
	*/
	af, bf := 5.67, 8.97
	fmt.Printf("\n\nfloat point \ntype of a : %T, b :  %T\n", af, bf)
	sum := af + bf
	diff := af - bf
	fmt.Println("sum :", sum, "diff :", diff)

	n1, n2 := 56, 89
	fmt.Println("sum :", n1+n2, "  diff :", n1-n2)

	/*
		complex types - complex64, complex128 - real and imaginary parts

		go has builtin function complex to construct complex no
		 - func complex(r, i FloatType) ComplexType
		 - takes a real and imaginary part as a parameter
		 - Both the real and imaginary parts must be of the same type.
		 - Both float32 -> complex64, float64 -> complex128
	*/
	c1 := complex(5, 7)
	c2 := 8 + 27i
	fmt.Println("\ncomplex number \nc1:", c1, "c2:", c2)
	cadd := c1 + c2
	fmt.Println("sum :", cadd)
	cmul := c1 * c2
	fmt.Println("mul :", cmul)

	// string are collection of bytes in Go
	first := "Darshan"
	last := "Sudake"
	name := first + " " + last
	fmt.Println("\nstring \nMy name is", name)

	/*
		type conversion
		 - Go is very strict about explicit typing
		 - no automatic type promotion or conversion
		 - T(v) : convert a value v to type T
	*/
	i := 55   // int
	j := 67.8 // float64
	// add := i + j // int + float64 not allowed in go
	add := i + int(j) // j is converted to int
	fmt.Println("\ntype conversion")
	fmt.Println(i, "+", j, "= ", add)

	// Explicit type conversion is required to assign a variable of one type to another
	k := 10
	var l float64 = float64(k) //this statement will not work without explicit conversion
	// try to assign i to j without any type conversion, the compiler will throw an error
	fmt.Println("l :", l)

}
