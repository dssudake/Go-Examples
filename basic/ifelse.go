// Program to learn conditional statements (if-else) in Golang

/*
	'if else' - if statement has a boolean condition
	 - It executes a block of code if that condition evaluates to true
	 - It executes an alternate else block if the condition evaluates to false

	syntax -->
	 - braces {} are mandatory even if there is only one line of code between braces{}
	 - 'else' construct will be executed if condition in the 'if' statement evaluates to false

		if condition {
		} else {
		}

	 - whichever if or else if's condition evaluates to true, it's corresponding code block is executed
	 - if none of the conditions are true then else block is executed
		if condition1 {
		...
		} else if condition2 {
		...
		} else {
		...
		}

	 - If with assignment - optional shorthand assignment statement is executed before condition is evaluated
		if assignment-statement; condition {
		}
*/
package main

import "fmt"

func main() {
	// If else statement
	num := 11
	if num%2 == 0 { //checks if number is even
		fmt.Println(num, "is an even number")
	} else {
		fmt.Println(num, "is an odd number")
	}

	// If ... else if ... else statement
	x := 99
	if x <= 50 {
		fmt.Println(x, "is less than or equal to 50")
	} else if x >= 51 && x <= 100 {
		fmt.Println(x, "is between 51 and 100")
	} else {
		fmt.Println(x, "is greater than 100")
	}

	// If with assignment
	// the scope of num is limited to if else blocks
	if y := 10; y%2 == 0 { //checks if number is even
		fmt.Println(y, "is an even number")
	} else {
		fmt.Println(y, "is an odd number")
	}

	/*
		else statement should start in same line after closing curly brace } of if statement
		because in the rules, it's specified that a semicolon will be inserted after
		closing brace }, if that is the final token of the line. Since if{...} else {...} is
		one single statement, a semicolon should not be present in the middle of it
		if num % 2 == 0 {
			...
		}
		else {
			...
		}

		Idiomatic Go -
		 In Go's philosophy, it is better to avoid unnecessary branches and indentation of code.
		 It is also considered better to return as early as possible.
	*/
	z := 10
	if z%2 == 0 { //checks if number is even
		fmt.Println(z, "is an even number")
		return
	}
	fmt.Println(z, "is an odd number")
}
