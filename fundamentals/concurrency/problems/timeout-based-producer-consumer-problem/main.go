package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
Timeout-Based Producer-Consumer Pattern:
Implement a producer-consumer pattern where the producer generates random numbers and sends them to a consumer via a channel. The consumer processes and prints these numbers. Introduce a timeout mechanism using a `time.After` function to stop the producer and consumer after a specified duration. Ensure that all channels are closed properly and that the timeout is handled gracefully.
**/

func producer(numbers chan<- int) {
	for {
		num := rand.Intn(1e3)
		fmt.Println("Produced: ", num)
		numbers <- num
		time.Sleep(1 * time.Second)
	}
}

func consumer(numbers <-chan int) {
	for num := range numbers {
		fmt.Println("Consumed: ", num)
	}
}

func main() {
	numbers := make(chan int)
	timeout := time.After(10 * time.Second)

	go producer(numbers)
	go consumer(numbers)

	select {
	case <-timeout:
		close(numbers)
		fmt.Println("Timeout reached, stopping the programs")

	}
}
