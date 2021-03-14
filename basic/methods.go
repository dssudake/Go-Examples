// Program to learn methods in Golang

/*
	method -
	 - a function with a special receiver type between func keyword and method name
	 - receiver can either be a struct type or non-struct type

	method declaration -->
		func (t Type) methodName(parameter list) {
		}

	why methods instead of functions --->
	 - Go is not a pure object-oriented programming language (no  classes)
	   Hence methods on types are a way to achieve behavior similar to classes. 
	   Methods allow logical grouping of behavior related to type similar to classes. 
	   In above sample program, all behaviors related to Employee type can be grouped by creating methods using Employee receiver type.
	   For example, we can add methods like calculatePension, calculateLeaves and so on.
	 - Methods with same name can be defined on different types whereas functions with same names are not allowed. 
	   Let's assume that we have a Square and Circle structure. It's possible to define a method named Area on both Square and Circle. 
*/
package main


import (  
	"fmt"
	"math"
)

type Employee struct {  
	name     string
	salary   int
	currency string
}

// displaySalary() method has Employee as the receiver type
func (e Employee) displaySalary() {  
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

type Rectangle struct {  
	length int
	width  int
}

type Circle struct {  
	radius float64
}

func (r Rectangle) Area() int {  
	return r.length * r.width
}

func (c Circle) Area() float64 {  
	return math.Pi * c.radius * c.radius
}

func main() {  
	emp1 := Employee {
			name:     "Sam Adolf",
			salary:   5000,
			currency: "$",
	}
	emp1.displaySalary() // Calling displaySalary() method of Employee type
	// displaySalary(emp1)  // same program with function

	r := Rectangle{
		length: 10,
		width:  5,
	}
	fmt.Println("\nArea of rectangle :", r.Area())
	
	c := Circle{
		radius: 12,
	}
	fmt.Println("Area of circle :", c.Area())
}
