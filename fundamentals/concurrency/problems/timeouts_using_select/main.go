package main

import (
	"fmt"
	"time"
)

/**
Create a program that simulates a task that may sometimes take too long to complete.
Use a `select` statement to implement a timeout mechanism.
If the task completes within a given time frame,
print the result; otherwise, print a timeout message.
**/

func simulateTask(duration time.Duration, resultChan chan<- string) {
	time.Sleep(duration)
	resultChan <- "Task completed"
}

func main() {
	resultChan := make(chan string)
	timeout := 2 * time.Second

	go simulateTask(3*time.Second, resultChan)

	select {
	case result := <-resultChan:
		fmt.Println(result)
	case <- time.After(timeout):
		fmt.Println("timeout: task took too long to finish")
	}
	fmt.Println("fin")
}
