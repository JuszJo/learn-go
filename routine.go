package main 

import (
	"fmt"
	// "time"
)
func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {
	tst := make(chan string)

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
		tst <- msg
	}("going")

	wait := <-tst

	// time.Sleep(time.Second)

	fmt.Println("Waited", wait)

	fmt.Println("Done")
}
