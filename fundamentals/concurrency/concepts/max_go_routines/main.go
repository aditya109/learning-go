package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"sync"
	"time"
)

type task struct {
	id  int
	msg string
}

const maxGoroutines = 5

func worker(id int, tasks chan task, wg *sync.WaitGroup) {
	defer wg.Done()
	for t := range tasks {
		fmt.Printf("worker %d received task %s\n", id, t.msg)
		time.Sleep(time.Duration(rand.IntN(1e4)) * time.Millisecond)
		fmt.Printf("worker %d finished task %s\n", id, t.msg)
	}
}

func main() {
	tasks := make(chan task, maxGoroutines)

	var wg sync.WaitGroup

	for i := 0; i < maxGoroutines; i++ {
		wg.Add(1)
		go worker(i, tasks, &wg)
	}

	for i := 0; i < 20; i++ {
		tasks <- task{
			id:  i,
			msg: "Target for task #" + strconv.Itoa(i),
		}
	}

	close(tasks)
}
