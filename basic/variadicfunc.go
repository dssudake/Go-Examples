// Program to learn variadic functions in Golang

/*
	variadic function -
	 - general function accept only a fixed number of arguments
	 - variadic function accepts a variable number of arguments
	 - If last parameter of function definition is prefixed by ellipsis ...
	   then function can accept any number of arguments for that parameter
	 - Only the last parameter of a function can be variadic

	function declaration -->
		// func hello(a int, b ...int) { // parameter b is variadic
		}

	 - function can be called by using syntax
	   hello(1, 2) //passing one argument "2" to b
	   hello(5, 6, 7, 8, 9) //passing arguments "6, 7, 8 and 9" to b

	 - It is also possible to pass zero arguments to a variadic function
	   hello(1)
*/
package main

import "fmt"

// variadic functions work by converting variable number of arguments to a slice of type of variadic parameter
func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func findSlice(num int, nums []int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}
func main() {
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 45, 109)
	find(78, 38, 56, 98)
	// not passed any argument to variadic nums ...int parameter.
	// in this case, nums will be a nil slice with length and capacity 0.
	find(87)

	/*
		Slice arguments vs Variadic arguments
		 - why do we even need variadic func when we can achieve same functionality using slices

		following of the advantages of using variadic arguments?
		 - There is no need to create a slice during each function call.
		 - empty slice need to be created just to satisfy the signature of the find function.
		   This is totally not needed in case of variadic functions.
		 - program with variadic functions is more readable than once with slices.
	*/
	findSlice(89, []int{89, 90, 95})
	findSlice(45, []int{56, 67, 45, 90, 109})
	findSlice(78, []int{38, 56, 98})
	findSlice(87, []int{})

	/*
		Append is a variadic function
		append function in the standard library used to append values to a slice accepts any number of arguments
		 func append(slice []Type, elems ...Type) []Type
	*/

	/*
		Passing a slice to a variadic function

		This will not work, Program will fail with compilation error
		 nums := []int{89, 90, 95}
		 find(89, nums)
	*/

	// there a way to pass a slice to a variadic function. we have to suffix the slice with ellipsis ...
	// If that is done, the slice is directly passed to the function without a new slice being created
	nums := []int{89, 90, 95}
	find(89, nums...)
}
