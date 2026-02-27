package main

import (
	"fmt"
	"time"
)

func main() {
	// using go syntax to enable concurrency

	done := make(chan bool)

	go greet("Rajesh", done)
	go slowgreet("Rajesh", done)
	go greet("Rajesh", done)

	for range done { // this method only works when you know, which function will finish the last, where you use close() with the last finishing func.
	}
}

func greet(greeting string, donechan chan bool) {
	fmt.Println(greeting, "there !")
	donechan <- true
}

func slowgreet(greeting string, doneChan chan bool) {
	time.Sleep(5 * time.Second) // simulate slow func
	fmt.Println("hello, there: ", greeting)
	doneChan <- true
	close(doneChan) // this method only works when you know, which function will finish the last.
}
