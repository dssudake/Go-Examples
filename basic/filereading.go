// Program to read files using Golang

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	// Reading an entire file into memory
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))

	// Using absolute file path
	data, err = ioutil.ReadFile("/home/chronin/go/src/Go-Examples/basic/test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))

	// Passing the file path as a command line flag
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()
	fmt.Println("value of fpath is", *fptr)
	data, err = ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}
