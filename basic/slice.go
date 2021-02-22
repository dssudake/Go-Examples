// Program to learn array in Golang

/*
	slices -
	 - convenient, flexible and powerful wrapper on top of array
	 - slices do not own any data on their own. They just references to existing arrays
	 - slice with elements of type T is represented by []T
*/
package main

import "fmt"

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}

func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	return countriesCpy
}

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	// syntax a[start:end] creates a slice from array a
	var b []int = a[1:4] // creates a slice from a[1] to a[3]
	fmt.Println("b :", b)

	c := []int{6, 7, 8} // creates an array and returns a slice reference
	fmt.Println("\nc :", c)

	/*
		Modifying a slice -
		 - A slice does not own any data of its own.
		 - It is just a representation of the underlying array.
		 - Any modifications done to slice will be reflected in underlying array
	*/
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before :", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after  :", darr)

	// When number of slices share same underlying array, changes that each one makes will be reflected in array
	numa := [3]int{78, 79, 80}
	nums1 := numa[:] // creates a slice which contains all elements of array
	nums2 := numa[:]
	fmt.Println("\nbefore change 1", numa)
	nums1[0] = 100
	fmt.Println("after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("after modification to slice nums2", numa)

	/*
		length and capacity of a slice -
		 length   - number of elements in the slice
		 capacity - no of elements in underlying array starting from index from which slice is created
	*/
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("\nlength of slice %d capacity %d", len(fruitslice), cap(fruitslice)) //length of fruitslice is 2 and capacity is 6

	// slice can be re-sliced upto its capacity. Anything beyond that will cause program to throw run time error
	fruitslice = fruitslice[:cap(fruitslice)] //re-slicing furitslice till its capacity
	fmt.Println("\nAfter re-slicing length is", len(fruitslice), "and capacity is", cap(fruitslice))

	/*
		creating a slice using make -
		 - func make([]T, len, cap) []T can be used to create slice by passing type, length & capacity
		 - capacity parameter is optional and defaults to length
		 - make function creates array and returns a slice reference to it
		 - values are zeroed by default when a slice is created using make
	*/
	i := make([]int, 5, 5)
	fmt.Println("\ni :", i)

	/*
		Appending to a slice -
		 - Slices are dynamic and new elements can be appended to slice using append function
		 - func append(s []T, x ...T) []T -->
		   x ...T - function accepts variable number of arguments for parameter x
			 These type of functions are called variadic functions

		If slices are backed by arrays which are of fixed length then how come a slice is of dynamic length?
		 - It's because when new elements are appended to slice, a new array is created
		 - elements of existing array are copied to new array and new slice reference for this new array is returned
		 - The capacity of the new slice is now twice that of the old slice
	*/
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("\nold -> length :", len(cars), ", capacity:", cap(cars), ", cars:", cars) //capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("new -> length :", len(cars), ", capacity:", cap(cars), ", cars:", cars) //capacity of cars is doubled to 6

	// zero value of a slice type is nil. It has has length and capacity 0
	// It is possible to append values to a nil slice using the append function
	var names []string //zero value of a slice is nil
	if names == nil {
		fmt.Println("\nslice is nil going to append")
		names = append(names, "John", "Sebastian", "Vinay")
		fmt.Println("names :", names)
	}

	// we can append one slice to another using the ... operator
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("\nfood :", food)

	/*
		Passing a slice to a function
		Slices can be thought of as being represented internally by a structure type
		type slice struct {
			Length        int
			Capacity      int
			ZerothElement *byte
		}

		A slice contains the length, capacity and a pointer to the zeroth element of the array.
		If slice is passed to function, even though it's passed by value, pointer variable will refer to same underlying array
		Hence when slice is passed to function as parameter, changes made inside function are visible outside function too
	*/
	nos := []int{8, 7, 6}
	fmt.Println("\nslice before function call :", nos)
	subtactOne(nos)                                  //function modifies the slice
	fmt.Println("slice after function call  :", nos) //modifications are visible outside
	fmt.Println()

	// Multidimensional slices - Similar to arrays, slices can have multiple dimensions
	pls := [][]string{
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}

	/*
		Memory Optimisation -
		 - Large array will still be in memory as slice still refernce small part of it
		 - As long as slice is in memory, array cannot be garbage collected
		 - To improve memory management we can use copy function 'func copy(dst, src []T) int' to make copy of slice
	*/
	fmt.Println("\nMemory optimization using copy")
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}
