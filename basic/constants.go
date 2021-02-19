// Program to learn constants in Golang

/*
	constant - denotes fixed values
	 - keyword const is used to declare a constant.
	 - value of constant should be known at compile time

	declare a single constant --> const name = value

	declare multiple constants -->
	const (
      name1 = initialvalue1
      name2 = initialvalue2
	)
*/
package main

import "fmt"

func main() {
	const num = 10 // const declaration
	fmt.Println("num :", num)

	/*
		Constants, as the name indicate, cannot be reassigned again to any other value
			const a = 55 //allowed
			a = 89 //reassignment not allowed

		Value of a constant should be known at compile time
			var a = math.Sqrt(4)   //allowed
		  const b = math.Sqrt(4) //not allowed
	*/
	// Declare group of constants
	const (
		name    = "Sam"
		age     = 29
		country = "canada"
	)
	fmt.Println("\nname :", name, ", age :", age, ", country :", country)

	/*
		String Constants, Typed and Untyped Constants
		 - Any value enclosed between double quotes is a string constant
		 - type of string constant --> untyped

		const hello = "Hello World"  // constant hello doesn't have a type.

		Go is a strongly typed language. All variables require an explicit type.

		untyped constants have a default type associated with them and they supply it if and only if a line of code demands it
	*/
	const n = "Sam"
	var str = n
	fmt.Printf("\ntype %T value %v", str, str)

	// typedhello in the below code is a constant of type string
	const typedhello string = "Hello World"

	// Mixing types during the assignment is not allowed
	var defaultName = "Sam"         // allowed
	type myString string            // new type myString which is an alias of string
	var customName myString = "Sam" // allowed
	// customName = defaultName     // not allowed string != myString
	fmt.Printf("\ntype -> defaultName : %T , customName:  %T", defaultName, customName)

	// Boolean constants are no different from string constants
	// They are two untyped constants true and false
	const trueConst = true
	type myBool bool
	var defaultBool = trueConst       //allowed
	var customBool myBool = trueConst //allowed
	// defaultBool = customBool       //not allowed bool != myBool
	fmt.Printf("\ntype -> defaultBool : %T , customBool:  %T\n", defaultBool, customBool)

	// Numeric Constants - type of each variable is determined by the syntax of the numeric constant
	var i = 5
	var f = 5.6
	var c = 5 + 6i
	fmt.Printf("\ntype i : %T, f : %T, c : %T\n", i, f, c)

	/*
		In example below syntax of a is generic. It can represent a float, int or even a complex no with
		no imaginary part. Hence it is possible to be assigned to any compatible type. The default type
		of these kinds of constants can be thought of as being generated on the fly depending on the context
	*/
	const a = 5
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println("\nintVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)

	/*
		Numeric Expression -
		Numeric constants are free to be mixed and matched in expressions and a type is needed only when
		they are assigned to variables or used in any place in code which demands a type.
	*/
	var x = 5.9 / 8
	fmt.Printf("\nx type : %T , value : %v", x, x)
}
