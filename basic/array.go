// Program to learn array in Golang

/*
	arrays -
	 - Collection of elements that belong to the same type
	 - Mixing values of different types is not allowed in Go
	 - All elements in array are automatically assigned zero value of array type
	 - Index of array starts from 0 and ends at (length - 1)

	array declaration --> [n]T
	 n : number of elements in array, T : type of each element
*/
package main

import "fmt"

func changeLocal(x [5]int) {
	x[0] = 55
	fmt.Println("inside function :", x)
}

func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func main() {
	var a [3]int //int array with length 3
	fmt.Println("a :", a)

	a[0] = 12 // array index starts at 0
	a[1] = 78
	a[2] = 50
	fmt.Println("\na :", a)

	b := [3]int{12, 78, 50} // short hand declaration to create array
	fmt.Println("\nb :", b)

	// It is not necessary to assign a value  to all elements in an array during short hand declaration
	// remaining elements are assigned 0 automatically
	c := [3]int{12}
	fmt.Println("\nc :", c)

	// We can ignore length of the array in the declaration and replace it with '...'
	// In that case compiler will find length
	d := [...]int{12, 78, 50} // ... makes the compiler determine the length
	fmt.Println("\nd :", d)

	/*
		The size of the array is a part of the type
		Hence [5]int and [25]int are distinct types. Because of this, arrays cannot be resized
		a := [3]int{5, 78, 8}
		var b [5]int
		b = a // not possible since [3]int and [5]int are distinct types
	*/

	/*
		Arrays are value types
		 - Arrays in Go are value types and not reference types
		 - Copy of original array is assigned to new variable
		 - If changes are made to new variable, it will not be reflected in original array
	*/
	ca := [...]string{"USA", "China", "India", "Germany", "France"}
	cb := ca // a copy of a is assigned to b
	cb[0] = "Singapore"
	fmt.Println("\nca :", ca)
	fmt.Println("cb :", cb)

	// Similarly when arrays are passed by value to functions as parameters then original array in unchanged
	x := [...]int{5, 6, 7, 8, 8}
	fmt.Println("\nbefore passing to function :", x)
	changeLocal(x) //x is passed by value
	fmt.Println("after passing to function :", x)

	// length of array is found by passing array as parameter to 'len' function
	la := [...]float64{67.7, 89.8, 21, 78}
	fmt.Println("\nlength of la :", len(la))

	for i := 0; i < len(la); i++ { //looping from 0 to the length of the array
		fmt.Printf("la element %d : %.2f\n", i, la[i])
	}
	fmt.Println()

	// range returns both index and value at that index
	sum := float64(0)
	for i, v := range la { // range returns both the index and value
		fmt.Printf("la element %d : %.2f\n", i, v)
		sum += v
	}
	fmt.Println("sum of all elements of la", sum)

	/*
		If you want only value and not index, then replace index with the _ blank identifier
		for _, v := range a { //ignores index
		}
		The above for loop ignores the index. Similarly the value can also be ignored.
	*/

	// Multidimensional arrays
	ma := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, // this comma is necessary. compiler will complain if you omit this comma
	}
	fmt.Println()
	printarray(ma)

	/*
		Array has restriction of being fixed length. It is not possible to increase length of array.
		This is were slices come into picture, slices are more common than conventional arrays.
	*/
}
