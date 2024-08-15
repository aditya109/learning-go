package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/**
**Error Handling and Recovery:**

Design a producer-consumer system where the producer generates numbers and sends them to the consumer.
The consumer processes these numbers but occasionally encounters errors (e.g., invalid data).
Introduce an error channel where the consumer can send errors.
Implement a mechanism to handle errors gracefully and stop both the producer and consumer when a critical error is encountered.

**/

type Result struct {
	num int
	err error
}

func producerWithWaitGroup(results chan<- Result, quit <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
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

func consumerWithWaitGroup(results <-chan Result, quit chan<- bool, err chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
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

func ExecuteWithWaitGroups() {
	var results = make(chan Result)
	var errChan = make(chan error)
	var quit = make(chan bool)

	var wg sync.WaitGroup

	wg.Add(1)
	go producerWithWaitGroup(results, quit, &wg)
	wg.Add(1)
	go consumerWithWaitGroup(results, quit, errChan, &wg)

	go func() {
		wg.Wait()
		close(errChan)
	}()

	if err, ok := <-errChan; ok {
		fmt.Printf("error faced, err : %v", err)
	} else {
		fmt.Println("No error encountered.")
	}

	fmt.Println()
}
