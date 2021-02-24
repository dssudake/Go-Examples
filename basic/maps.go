// Program to learn maps in Golang

/*
	maps -
	 - a builtin type in Go that is used to store key-value pairs
	 - Maps are similar to dictionaries in Python
	 - created by passing the type of key and value to the make function
	   make(map[type of key]type of value)
*/
package main

import "fmt"

type employee struct {
	salary  int
	country string
}

func main() {
	employeeSalary := make(map[string]int) // creates a map named employeeSalary with string key and int value
	fmt.Println(employeeSalary)            // map[]

	employeeSalary["steve"] = 12000
	employeeSalary["jamie"] = 15000
	employeeSalary["mike"] = 9000
	fmt.Println("employeeSalary map contents  :", employeeSalary)

	// possible to initialize a map during the declaration itself
	employeeSalary2 := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	employeeSalary2["mike"] = 9000
	fmt.Println("employeeSalary2 map contents :", employeeSalary2)

	// All comparable types such as boolean, integer, float, complex, string, ... can also be keys.
	// Even user-defined types such as structs can be keys

	/*
		Zero value of a map -
		 The zero value of a map is nil. If you try to add elements to a nil map, a run-time panic will occur.
		 Hence the map has to be initialized before adding elements

		 var employeeSalary map[string]int
		 employeeSalary["steve"] = 12000
	*/

	// Retrieving value for a key from a map
	employeeName := "jamie"
	salary := employeeSalary[employeeName]
	fmt.Println("\nSalary of", employeeName, ":", salary)

	// if we try to access element which is not present, zero value of type of that element will be returned.
	fmt.Println("\nSalary of joe :", employeeSalary["joe"])

	/*
		checking if key exists -
		 value, ok := map[key]
		 if ok is true, then the key is present and its value is present in variable value,
		 else the key is absent
	*/
	newEmp := "joe"
	value, ok := employeeSalary[newEmp]
	if ok == true {
		fmt.Println("Salary of", newEmp, "is", value)
		return
	}
	fmt.Println(newEmp, "not found")

	/*
		Iterate over all elements in a map -
		 order of retrieval of values from map when using for range is not guaranteed to be same for each execution of program.
		 It is also not the same as the order in which the elements were added to the map
	*/
	fmt.Println("\nContents of the map")
	for key, value := range employeeSalary {
		fmt.Printf("employeeSalary[%s] = %d\n", key, value)
	}

	/*
		Deleting items from a map -
		 delete(map, key) is the syntax to delete key from a map
		 The delete function does not return any value
		 If we try to delete a key that is not present in map, there will be no runtime error.
	*/
	delete(employeeSalary, "steve")
	fmt.Println("\nmap after deletion :", employeeSalary)

	/*
		Map of structs -
	*/
	emp1 := employee{
		salary:  12000,
		country: "USA",
	}
	emp2 := employee{
		salary:  14000,
		country: "Canada",
	}
	emp3 := employee{
		salary:  13000,
		country: "India",
	}
	employeeInfo := map[string]employee{
		"Steve": emp1,
		"Jamie": emp2,
		"Mike":  emp3,
	}

	fmt.Println()
	for name, info := range employeeInfo {
		fmt.Printf("Employee: %s Salary:$%d  Country: %s\n", name, info.salary, info.country)
	}

	// Length of the map
	fmt.Println("\nlength is employeeInfo :", len(employeeInfo))

	/*
		Maps are reference types -
		 Similar to slices, maps are reference types. When map is assigned to new variable,
		 they both point to same internal data structure

		 when maps are passed as parameters to functions, any changes made to map inside function,
		 it will be visible to caller also
	*/
	fmt.Println("\nOriginal employee salary :", employeeSalary)
	modified := employeeSalary
	modified["mike"] = 18000
	fmt.Println("Employee salary changed  :", employeeSalary)

	/*
		Maps equality -
		Maps can't be compared using the == operator. The == can be only used to check if a map is nil
		One way to check whether two maps are equal is to compare each one's individual elements one by one.
		The other way is using reflection.
	*/
}
