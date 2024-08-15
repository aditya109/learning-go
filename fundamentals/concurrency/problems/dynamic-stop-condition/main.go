package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/**
**Dynamic Stop Condition:**
Create a producer-consumer pattern where the producer generates integers and sends them to a consumer.
The consumer should process and print the numbers.
Introduce a stop condition where the producer stops producing and signals the consumer to stop if the sum of generated numbers exceeds a specified threshold. Use a stop channel to communicate this condition and ensure proper closure of all channels.
**/

func producer(numbers chan<- int, quit <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		num := rand.Intn(1e3)
		fmt.Println("Produced: ", num)
		select {
		case numbers <- num:
		case <-quit:
			fmt.Println("Producer received stop command during sending...")
			close(numbers) // Ensure the channel is closed properly
			return
		}
	}
}

func consumer(target int, numbers <-chan int, quit chan<- bool, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for num := range numbers {
		sum += num
		if sum > target {
			fmt.Println("accumulated sum: ", sum)
			quit <- true
			close(quit)
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	numbers, quit := make(chan int), make(chan bool)
	target := rand.Intn(1e3) * 50
	fmt.Println("target = ", target)
	wg.Add(2)
	go producer(numbers, quit, &wg)
	// wg.Add(1)
	go consumer(target, numbers, quit, &wg)

	wg.Wait()
	fmt.Println("Fin")
}
