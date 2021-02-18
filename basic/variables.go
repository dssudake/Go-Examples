// Program to learn variables in Golang

/*
	Variable - Name given to memory location to store value of specific type

	declare a single variable --> var name type

	declare a multiple variables --> var name1, name2 type = initialvalue1, initialvalue2

	declare variables belonging to different types -->
	var (
      name1 = initialvalue1
      name2 = initialvalue2
	)

	short hand declaration --> name := initialvalue
*/
package main

import "fmt"

func main() {
	var num int // variable declaration

	// If a variable is not assigned any value,
	// Go automatically initializes it with the zero value of the variable's type.
	fmt.Println("num (default) :", num)

	num = 10 // assignment
	fmt.Println("num (updated) :", num)

	// Go will automatically be able to infer the type of that variable using that initial value
	var inum = 24
	fmt.Println("\ninum (initial) :", inum)

	// Multiple variable declaration
	// var width, height int = 100, 50 // declaring multiple variables
	var length, breadth = 100, 50 // "int" can be dropped
	fmt.Println("\nlength is", length, "breadth is", breadth)

	// Multiple variable declaration - different types
	var (
		name   = "Sam"
		age    = 29
		height int
	)
	fmt.Println("\nmy name is", name, ", age is", age, "and height is", height)

	// Short Hand Declaraton(SHD)
	count := 10
	fmt.Println("\ncount :", count)

}

/*
	<--- More about Short Hand Declaration -->

	name, age := "naveen", 29 //SHD

	SHD requires initial values for all variables on the left-hand side of the assignment
		name, age := "naveen" //error

	SHD syntax can be used only when at least one of the variables on the left side of := is newly declared
  	a, b := 20, 30 // declare variables a and b
   	b, c := 40, 50 // b is already declared but c is new
   	b, c = 80, 90 // assign new values to already declared variables b and c

   	a, b := 20, 30 //a and b declared
   	a, b := 40, 50 //error, no new variables

  Variables can also be assigned values which are computed during run time
		a, b := 145.8, 543.8
   	c := math.Min(a, b)

  Go is strongly typed, variables declared as belonging to one type cannot be assigned a value of another type
		age := 29      // age is int
    age = "naveen" // error since we are trying to assign a string to a variable of type int
*/
