package main

/**
You are tasked with processing a large dataset of integers. There are three stages in the pipeline:
1. **Generate**: A goroutine generates a sequence of integers.
2. **Square**: Multiple goroutines take integers from the input channel, square them, and send them to an output channel.
3. **Sum**: A goroutine sums the squared integers from the output channel and sends the final sum to a result channel.

Implement this using fan-in and fan-out patterns with buffered channels.
**/

func generate(lowerBound, upperBound int) <-chan int {
	integers := make(chan int)
	go func() {
		defer close(integers)
		for i := lowerBound; i <= upperBound; i++ {
			integers <- lowerBound + upperBound
		}
	}()
	return integers
}

func square(input <-chan int) <-chan int {
	squares := make(chan int)
	go func() {
		defer close(squares)
		for value := range input {
			squares <- value * value
		}
	}()
	return squares
}
