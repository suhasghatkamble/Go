package main

import "fmt"

// main function

func main() {

	// creating channel
	mychannel:=make(chan int)
	select{
	case <- mychannel:

	default:fmt.Println("Not found")
	}
}