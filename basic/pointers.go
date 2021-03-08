// Program to learn pointers in Golang

/*
	Pointers :
	 - A pointer is variable which stores memory address of another variable
	 - *T is type of pointer variable which points to value of type T
	 - zero value of a pointer is nil
*/
package main

import "fmt"

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
}
