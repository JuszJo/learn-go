package main

import "fmt"

func main() {
	channel := make(chan string, 2)

	channel <- "first"
	channel <- "second"

	message := <-channel
	buffer := <-channel

	fmt.Println(message)
	fmt.Println(buffer)
}