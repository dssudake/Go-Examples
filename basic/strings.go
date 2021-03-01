// Program to learn strings in Golang

/*
	Strings :
	 - Different in implementation when compared to other languages
	 - It is a slice of bytes in Go
	 - Strings can be created by enclosing set of characters inside double quotes " "
	 - Strings in Go are Unicode compliant and are UTF-8 Encoded.
*/
package main

import "fmt"

// Accessing individual bytes of a string
func printBytes(s string) {
	fmt.Printf("Bytes: ")
	// len(s) returns number of bytes in the string
	for i := 0; i < len(s); i++ {
		// %x is the format specifier for hexadecimal
		fmt.Printf("%x ", s[i])
	}
}

// Accessing individual characters of a string
func printChars(s string) {
	fmt.Printf("Characters: ")
	for i := 0; i < len(s); i++ {
		// %c format specifier is used to print characters of the string
		fmt.Printf("%c ", s[i])
	}
}

func main() {
	name := "Hello World"
	fmt.Println(name)

	// %s is format specifier to print a string
	fmt.Printf("String: %s\n", name)

	// Accessing individual bytes of a string
	printBytes(name)
	fmt.Println()

	// Accessing individual characters of a string
	printChars(name)
}
