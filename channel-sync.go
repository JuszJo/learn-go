package main

import (
	"fmt"
	"time"
)

func worker(chan done bool) {
	fmt.Println("Working...")
	time.Sleep(time.Second)
	fmt.Println("Done")

	done <- true
}

func main() {
	
}