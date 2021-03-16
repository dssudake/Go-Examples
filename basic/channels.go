// Program to learn channels in Golang

/*
	Channels -
	 - pipes using which Goroutines communicate
	 - Each channel has type associated with it i.e. type of data that channel is allowed to transport
	 - No other type is allowed to be transported using channel
	 - zero value of a channel is nil
	 - nil channels are not of any use, hence channel has to be defined using make similar to maps & slices

	 - direction of the arrow with respect to the channel specifies whether the data is sent or received
	    data := <- a // read from channel a
	    a <- data // write to channel a

	 - Sends and receives are blocking by default
	   When data is sent to a channel, control is blocked in send statement until some other Goroutine reads from channel
		 Similarly, when data is read from a channel, read is blocked until some Goroutine writes data to that channel.
*/
package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func main() {
	var a chan int
	// a := make(chan int)  --> short hand declaration is also a valid
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T\n\n", a)
	}

	done := make(chan bool)
	go hello(done)
	<-done
	fmt.Println("main function")

	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	// main Goroutine waits for data from both these channels
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("\nFinal output", squares+cubes)

	/*
		Deadlock - program will panic at runtime with Deadlock,
		 if one Goroutine sends data on a channel and no other Goroutine receives it and vice versa

		ch := make(chan int)
		ch <- 5
	*/

	/*
		Unidirectional channels --> channels that only send or receive data
		 It is possible to convert a bidirectional channel to a send only or receive only channel but not vice versa.
	*/

	/*
		Closing channels and for range loops on channels
		Senders have ability to close channel to notify receivers that no more data will be sent on channel.
		Receivers can use additional variable while receiving data from channel to check whether channel has been closed
		 v, ok := <- ch

		 ok : true --> value was received by a successful send operation to a channel
		    : false --> reading from a closed channel and value will be zero value of channel's type
	*/
	fmt.Println()
	ch := make(chan int)
	go producer(ch)
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}

	// using a for range loop
	fmt.Println()
	ch2 := make(chan int)
	go producer(ch2)
	for v := range ch2 {
		fmt.Println("Received ", v)
	}
}
