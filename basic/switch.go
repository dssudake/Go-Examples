// Program to learn conditional statement (switch) in Golang

/*
	'switch' -
	 - conditional statement that evaluates an expression
	 - compares it against a list of possible matches and executes the corresponding block of code
	 - It can be considered as an idiomatic way of replacing complex if else clauses
	 - Duplicate cases with the same constant value are not allowed
	 - Default case will be executed when none of other cases match
	 - Default case can be present anywhere in switch.
*/
package main

import (
	"fmt"
	"math/rand"
)

func number() int {
	num := 15 * 5
	return num
}

func main() {
	finger := 3
	fmt.Printf("Finger %d is ", finger)
	// A switch can include an optional statement that is executed before the expression is evaluated
	// switch finger := 3; finger {
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	// case 4: //duplicate case not allowed
	// 	fmt.Println("Another Ring")
	case 5:
		fmt.Println("Pinky")
	default: //default case
		fmt.Println("incorrect finger number")
	}
	fmt.Println()

	// Multiple expressions in case
	// It is possible to include multiple expressions in case by separating them with comma.
	letter := "i"
	fmt.Printf("Letter %s is a ", letter)
	switch letter {
	case "a", "e", "i", "o", "u": //multiple expressions in case
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}
	fmt.Println()

	/*
		Expressionless switch -
		 - expression in a switch is optional and it can be omitted
		 - if omitted switch is considered to be switch true and each of case expression
		   is evaluated for truth and corresponding block of code is executed
		 - This type of switch can be considered as alternative to multiple if else clauses.
	*/
	num := 75
	switch { // expression is omitted
	case num >= 0 && num <= 50:
		fmt.Println(num, "is greater than 0 and less than 50")
	case num >= 51 && num <= 100:
		fmt.Println(num, "is greater than 51 and less than 100")
	case num >= 101:
		fmt.Println(num, "is greater than 100")
	}
	fmt.Println()

	/*
		Fallthrough -
		 - In Go, control comes out of switch statement immediately after case is executed
		 - fallthrough statement is used to transfer control to the first statement of the case
		   that is present immediately after the case which has been executed
		 - fallthrough should be the last statement in a case
		 - If it is present somewhere in middle, compiler will complain that fallthrough statement out of place
	*/
	// Switch and case expressions need not be only constants.
	switch n := number(); { //n is not a constant
	case n < 50:
		fmt.Println(n, "is lesser than 50")
		fallthrough
	case n < 100:
		fmt.Println(n, "is lesser than 100")
		fallthrough
	case n < 200:
		fmt.Println(n, "is lesser than 200")
	}
	fmt.Println()

	// Fallthrough happens even when case evaluates to false
	// fallthrough cannot be used in last case of switch since there are no more cases to fallthrough
	switch num := 25; {
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num > 100:
		fmt.Printf("%d is greater than 100\n", num)
	}
	fmt.Println()

	/*
		Breaking switch -
		 - break statement can be used to terminate a switch early before it completes
	*/
	switch num := -5; {
	case num < 50:
		if num < 0 {
			fmt.Println("break switch")
			break
		}
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
	}
	fmt.Println()

	/*
		Breaking the outer for loop -
		 If break statement is used without label, switch statement will only be broken
		 and loop will continue running. So labeling loop and using it in break statement
		 inside the switch is necessary to break the outer for loop
	*/
randloop:
	for {
		switch i := rand.Intn(100); {
		case i%2 == 0:
			fmt.Println("Generated even number is", i)
			break randloop
		}
	}
}
