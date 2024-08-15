package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
**Error Handling and Recovery:**

Design a producer-consumer system where the producer generates numbers and sends them to the consumer.
The consumer processes these numbers but occasionally encounters errors (e.g., invalid data).
Introduce an error channel where the consumer can send errors.
Implement a mechanism to handle errors gracefully and stop both the producer and consumer when a critical error is encountered.

**/

func producerWithoutWaitGroup(results chan<- Result, quit <-chan bool) {
	for {

		select {
		case <-quit:
			fmt.Println("received quit command")
			close(results)
			return
		default:
			var result Result
			chance := rand.Intn(1e6)
			if chance%9 != 0 {
				num := rand.Intn(1e3)
				fmt.Printf("Produced: num= %d\n", num)
				result.num = num
			} else {
				err := fmt.Errorf("some error")
				fmt.Printf("Produced: err= %s\n", err)
				result.err = err
			}
			results <- result
		}
	}
}

func consumerWithoutWaitGroup(results <-chan Result, quit chan<- bool, err chan<- error) {
	for result := range results {
		if result.err != nil {
			fmt.Println("err encountered, sending stop signal", result.err.Error())
			err <- result.err
			quit <- true
			return
		}
		fmt.Println("consumed: num: ", result.num)
	}
}

func ExecuteWithoutWaitGroups() {
	var results = make(chan Result)
	var errChan = make(chan error)
	var quit = make(chan bool)

	go producerWithoutWaitGroup(results, quit)
	go consumerWithoutWaitGroup(results, quit, errChan)

	select {
	case err, ok := <-errChan:
		if ok {
			fmt.Printf("error faced, err : %v", err)
		}
	case <-time.After(10 * time.Second):
		fmt.Println("No error encountered.")
	}

	close(quit)

	for res := range results {
		fmt.Println(res.num)
	}

	fmt.Println()
}
