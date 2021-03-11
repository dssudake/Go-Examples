// Program to learn structs in Golang

/*
	struct -
	 - A struct is a user-defined type that represents a collection of fields.
	 - used to group data into single unit rather than having each of them as separate values.
	 -
*/
package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
	// fields belonging to same type can be declared in a single line followed by the type name
	// firstName, lastName string
}

type PersonAF struct {
	string
	int
}

type Address struct {
	city  string
	state string
}

type Person struct {
	name    string
	age     int
	address Address
}

type PersonPF struct {
	name string
	age  int
	Address
}

type name struct {
	firstName string
	lastName  string
}

func main() {
	// creating struct specifying field names
	emp1 := Employee{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}
	fmt.Println("Employee 1 :", emp1)

	// creating struct without specifying field names
	// necessary to maintain order of fields to be same as specified in struct declaration
	emp2 := Employee{"Thomas", "Paul", 29, 800}
	fmt.Println("Employee 2 :", emp2)

	// Anonymous structs : Declare structs without creating a new data type
	// It only creates new struct variable emp3 and doesn't define any new struct type like named structs
	emp3 := struct {
		firstName string
		lastName  string
		age       int
		salary    int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}
	fmt.Println("Employee 3 :", emp3)

	// The dot . operator is used to access the individual fields of a struct
	fmt.Println("\nFirst Name:", emp1.firstName)
	fmt.Println("Last Name:", emp1.lastName)
	fmt.Println("Age:", emp1.age)
	fmt.Printf("Salary: $%d\n", emp1.salary)
	emp1.salary = 900
	fmt.Printf("New Salary: $%d\n", emp1.salary)

	// When a struct is defined and it is not explicitly initialized with any value,
	// the fields of the struct are assigned their zero values by default.
	var emp4 Employee //zero valued struct
	fmt.Println("\nEmployee 4 :", emp4)
	fmt.Println("First Name:", emp4.firstName)
	fmt.Println("Last Name:", emp4.lastName)
	fmt.Println("Age:", emp4.age)
	fmt.Println("Salary:", emp4.salary)

	// Pointers to a struct
	emp5 := &Employee{
		firstName: "Sam",
		lastName:  "Anderson",
		age:       55,
		salary:    6000,
	}
	fmt.Println("\nFirst Name:", (*emp5).firstName)
	// emp5.age can be used instead of the explicit dereference (*emp5).age
	fmt.Println("Age:", emp5.age, "\n")

	// Anonymous fields : structs with fields that contain only a type without the field name
	// by default the name of an anonymous field is the name of its type
	p1 := PersonAF{
		string: "naveen",
		int:    50,
	}
	fmt.Println(p1.string)
	fmt.Println(p1.int)

	// Nested structs : struct contains a field which in turn is a struct
	p2 := Person{
		name: "Naveen",
		age:  50,
		address: Address{
			city:  "Chicago",
			state: "Illinois",
		},
	}
	fmt.Println("\nName:", p2.name)
	fmt.Println("Age:", p2.age)
	fmt.Println("City:", p2.address.city)
	fmt.Println("State:", p2.address.state)

	// Promoted fields : Fields that belong to an anonymous struct field in a struct
	// they can be accessed as if they belong to struct which holds anonymous struct field
	p3 := PersonPF{
		name: "Naveen",
		age:  50,
		Address: Address{
			city:  "Chicago",
			state: "Illinois",
		},
	}
	fmt.Println("\nName:", p3.name)
	fmt.Println("Age:", p3.age)
	fmt.Println("City:", p3.city)         //city is promoted field
	fmt.Println("State:", p3.state, "\n") //state is promoted field

	/*
		Exported structs and fields :
		If a struct type starts with a capital letter, then it is an exported type and it can be accessed from other packages
		Similarly, if the fields of a struct start with caps, they can be accessed from other packages.
	*/

	/*
		Structs Equality :
		Structs are value types and are comparable if each of their fields are comparable.
		Two struct variables are considered equal if their corresponding fields are equal.

		Struct variables are not comparable if they contain fields that are not comparable
		e.g :
		type image struct {
		 	data map[int]int
		}
	*/
	name1 := name{
		firstName: "Steve",
		lastName:  "Jobs",
	}
	name2 := name{
		firstName: "Steve",
		lastName:  "Jobs",
	}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}

	name3 := name{
		firstName: "Steve",
		lastName:  "Jobs",
	}
	name4 := name{
		firstName: "Steve",
	}

	if name3 == name4 {
		fmt.Println("name3 and name4 are equal")
	} else {
		fmt.Println("name3 and name4 are not equal")
	}

}
