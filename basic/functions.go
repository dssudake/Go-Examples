// Program to learn functions in Golang

/*
	function -
	 - a block of code that performs a specific task
	 - takes input, performs some calculations on input & generates output

	function declaration -->
		func functionname(parametername type) returntype {
 			//function body
		}

	parameters and return type are optional in a function
	following syntax is also a valid function declaration.
		func functionname() { }
*/
package main

import "fmt"

// If consecutive parameters are of same type, it is enough to be written once at the end
// func calculateBill(price int, no int) int {
func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}

/*
	Multiple return values -
		It is possible to return multiple values from a function
		In that case they must be specified between ( and )
*/
func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

/*
	Named return values -
		It is possible to return named values from a function
		considered as being declared as a variable in the first line of the function

	Note : return statement does not explicitly return any value,
		they are automatically returned when return statement is encountered
*/
func rectPropsNamed(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return //no explicit return value
}

func main() {
	price, no := 90, 6
	totalPrice := calculateBill(price, no) // function call
	fmt.Println("Total Price :", totalPrice)

	// Multiple return values
	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("\nArea : %f , Perimeter : %f", area, perimeter)

	// Named return values
	areaN, perimeterN := rectPropsNamed(10.8, 5.6)
	fmt.Printf("\nArea : %f , Perimeter : %f", areaN, perimeterN)

	/*
		Blank Identifier -->  _
		It can be used in place of any value of any type

		e.g.: What if we only need the area and want to discard the perimeter
	*/
	area, _ = rectProps(10.8, 5.6) // perimeter is discarded
	fmt.Printf("\nArea : %f ", area)
}
