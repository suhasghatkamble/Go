package main

import "fmt"

func myfunc(ch chan int) {
	fmt.Println(234 + <-ch)
}

func myfunc2(mychnl chan string) {
	for i := 0; i < 4; i++ {
		mychnl <- "CDAC ACTS"
	}
	close(mychnl)
}

func main() {
	fmt.Print("start Main Method\n")

	// Creating a channel
	ch := make(chan int)

	go myfunc(ch)
	ch <- 23

	// creating a channel
	c := make(chan string)

	// calling Goroutine
	go myfunc2(c)

	// When the value of ok is 
	// set to true means the 	
	// channel is open and it 
	// can send or receive data
	// when the value of ok is set to 
	// false means the channel is closed

	for {
		res, ok := <-c
		if !ok {
			fmt.Println("Channel is closed", ok)
			break
		}
		fmt.Println("Channel Open ", res, ok)
	}
	fmt.Println("End of Main Method")
}