// Program to learn pointers in Golang

/*
	Pointers :
	 - A pointer is variable which stores memory address of another variable
	 - *T is type of pointer variable which points to value of type T
	 - zero value of a pointer is nil
*/
package main

import "fmt"

func change(val *int) {
	*val = 55
}

func hello() *int {
	i := 5
	return &i
}

func modify1(arr *[3]int) {
	(*arr)[0] = 90
	// a[x] is shorthand for (*a)[x]. So (*arr)[0] in above program can be replaced by arr[0]
	// arr[0] = 90
}

func modify(sls []int) {
	sls[0] = 90
}

func main() {
	b := 255
	var a *int = &b // & operator is used to get address of a variable
	fmt.Printf("Type of a is %T\n", a)
	// we will get different address for b since location of b can be anywhere in memory
	fmt.Println("address of b is", a)

	// Zero value of a pointer
	var c *int
	if c == nil {
		fmt.Println("\nc is", c)
		c = &b
		fmt.Println("c after initialization is", c)
	}

	/*
		Pointers can be created using new function
		new function takes type as argument and returns a pointer to
		a newly allocated zero value of type passed as argument
	*/
	size := new(int)
	fmt.Printf("\nSize value is %d, type is %T, address is %v\n", *size, size, size)
	*size = 85
	fmt.Println("New size value is", *size)

	// Dereferencing a pointer means accessing the value of the variable which the pointer points to.
	// *a is the syntax to deference a
	x := 255
	y := &x
	fmt.Println("\naddress of x is", y)
	fmt.Println("value of x is", *y)
	*y++
	fmt.Println("new value of x is", x)

	// Passing pointer to a function
	f1 := 58
	fmt.Println("\nvalue of f1 before function call is", f1)
	f2 := &f1
	change(f2)
	fmt.Println("value of f1 after function call is", f1)

	/*
		Returning pointer from a function -
		 It is perfectly legal for a function to return a pointer of a local variable.
		 The Go compiler is intelligent enough and it will allocate this variable on the heap.
		 The behavior of this code is undefined in programming languages such as C and C++ as
		 variable i goes out of scope once function hello returns. But in case of Go,
		 compiler does escape analysis and allocates i on heap as address escapes local scope
	*/
	d := hello()
	fmt.Println("\nValue of d", *d, "\n")

	// Do not pass pointer to an array as argument to function. Use slice instead.
	m1 := [3]int{89, 90, 91}
	modify1(&m1)
	fmt.Println(m1, "\n")

	// Although this way of passing a pointer to array as argument to function and making
	// modification to it works, it is not the idiomatic way of achieving this in Go
	m2 := [3]int{89, 90, 91}
	modify(m2[:])
	fmt.Println(m2)

	/*
		Go does not support pointer arithmetic
		b := [...]int{109, 110, 111}
		p := &b
		p++
	*/
}
