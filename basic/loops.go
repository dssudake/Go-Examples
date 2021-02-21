// Program to learn loops in Golang

/*
	loops -
	 - It is used to execute a block of code repeatedly
	 - for is the only loop available in Go
	 - Go doesn't have while or do while loops like c

	for loop syntax -->
		for initialisation; condition; post {
		}
	 - initialisation statement will be executed only once.
	 - if condition evaluates to true, body inside { } will be executed followed by post statement
	 - post statement will be executed after each successful iteration of the loop
	 - all 3 components initialisation, condition and post are optional in Go
*/
package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println("\n")

	/*
		break statement
		 - is used to terminate for loop abruptly before it finishes its
		 - normal execution and move control to the line of code just after the for loop
	*/
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break //loop is terminated if i > 5
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println("\nline after for loop\n")

	/*
		continue statement
		 - is used to skip current iteration of for loop
		 - All code present in for loop after continue statement won't be executed for current iteration
		 - The loop will move on to the next iteration
	*/
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i) // print all odd numbers from 1 to 10
	}
	fmt.Println("\n")

	// Nested for loops
	n := 5
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()

	// labels - used to break the outer loop from inside the inner for loop
outer:
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == j {
				break outer // we wan't to stop printing when i and j are equal
			}
		}
	}
	fmt.Println()

	// More examples - variations of for loop
	// This format can be considered as alternative for while loop
	i := 0
	for i <= 10 { // initialisation and post are omitted
		fmt.Printf("%d ", i)
		i += 2
	}
	fmt.Println("\n")

	// It is possible to declare and operate on multiple variables in for loop
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { //multiple initialisation and increment
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}

	/*
		infinite loop -
		for {
			fmt.Println("This is infinite loop")
		}
	*/
}
