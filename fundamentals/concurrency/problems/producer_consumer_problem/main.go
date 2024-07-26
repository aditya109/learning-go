package main

import (
	"fmt"
	"time"
)

/***
Implement a producer-consumer pattern where the producer generates numbers and sends them to a consumer via a channel. The consumer should process these numbers and print them. Introduce a quit channel that can signal the producer and consumer to stop when a certain condition (e.g., reaching a specific number) is met. Ensure that all channels are properly closed, and use `range` to receive values.
**/

func producer(numbers chan<- int, quit <-chan bool) {
	i := 0
	for {
		select {
		case numbers <- i:
			i++
			time.Sleep(1 * time.Second)
		case <-quit:
			fmt.Println("producer stopping....")
			close(numbers)
			return
		}
	}
}

func consumer(numbers <-chan int, quit chan<- bool) {
	for num := range numbers {
		fmt.Printf("consumed: %d\n", num)
		time.Sleep(1 * time.Second)
		if num >= 10 {
			fmt.Println("consumer reached the stop condition")
			quit <- true
			return
		}
	}
}

func main() {
	numbers := make(chan int)
	quit := make(chan bool)

	go producer(numbers, quit)
	go consumer(numbers, quit)

	time.Sleep(2 * time.Second)
	quit <- true
	fmt.Println("Fin")
}
