// Program to learn goroutine in Golang

/*
	Goroutine -
	 - functions or methods that run concurrently with other functions / methods
	 - light weight threads, cost of creating a Goroutine is tiny when compared to a thread
*/
package main

import (
	"fmt"
	"time"
)

func hello() {  
	fmt.Println("Hello world goroutine")
}

func numbers() {  
	for i := 1; i <= 5; i++ {
			time.Sleep(250 * time.Millisecond)
			fmt.Printf("%d ", i)
	}
}
func alphabets() {  
	for i := 'a'; i <= 'e'; i++ {
			time.Sleep(400 * time.Millisecond)
			fmt.Printf("%c ", i)
	}
}

func main() {  
	go hello()
	fmt.Println("main function")
	
	/* 
	 - When a new Goroutine is started, goroutine call returns immediately. Unlike functions, 
	   control does not wait for Goroutine to finish executing. The control returns immediately to
	   next line of code after the Goroutine call and any return values from Goroutine are ignored.
	 - The main Goroutine should be running for any other Goroutines to run. If the main Goroutine 
	   terminates then the program will be terminated and no other Goroutine will run. 
	*/
	go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("main function")

	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
}