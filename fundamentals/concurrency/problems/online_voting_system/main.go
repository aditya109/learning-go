package main

import (
	"fmt"
	"sync"
)

/**
Let's create a scenario for a problem where multiple goroutines need to update a shared counter concurrently.
The problem will be framed in the context of a simulated online voting system.

Scenario: Online Voting System

In an online voting system, users can vote for their favorite candidate.
The system records each vote and updates a shared counter representing the total votes each candidate has received.
To ensure the integrity of the voting process, the system must prevent race conditions where multiple users might update the vote count simultaneously, potentially leading to incorrect results.

Problem Statement

Write a Go program to simulate an online voting system where multiple goroutines represent users voting for a candidate.
The program should:

Shared Counter: Use a shared integer counter to represent the total number of votes.
Concurrency: Launch a number of goroutines to simulate users voting concurrently.
Synchronization: Use a sync.Mutex to protect the shared counter from concurrent writes.
Completion: Use a sync.WaitGroup to ensure that all goroutines complete before printing the final vote count.
Output: Print the final value of the counter, which should match the total number of simulated votes.
Constraints and Assumptions
Assume that all votes are for a single candidate, so there is only one shared counter.
The number of votes (i.e., the number of goroutines) is a predefined constant.
Each goroutine increments the counter exactly once.
**/

type SharedCounter struct {
	mutex   *sync.Mutex
	counter int
}

func (c *SharedCounter) PerformVotingAction(data interface{}, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		c.mutex.Lock()
		defer c.mutex.Unlock()
		c.counter++
	}()
}

func main() {
	const numVotes = 1000
	c := &SharedCounter{
		counter: 0,
		mutex:   &sync.Mutex{},
	}
	var wg sync.WaitGroup
	for i := 0; i < numVotes; i++ {
		wg.Add(1)
		c.PerformVotingAction(nil, &wg)
	}
	wg.Wait()
	fmt.Printf("Final vote count: %d\n", c.counter)
}
