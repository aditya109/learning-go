package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Task func()

type Worker struct {
	id       int
	taskCh   chan Task
	taskWg   *sync.WaitGroup
	stealChs []chan Task
}

func (w *Worker) run() {
	for {
		select {
		case task := <-w.taskCh:
			task()
			w.taskWg.Done()
		default:
			w.stealTask()
		}
	}
}

func (w *Worker) stealTask() {
	for _, ch := range w.stealChs {
		select {
		case stolenTask := <-ch:
			stolenTask()
			w.taskWg.Done()
		default:
			time.Sleep(10 * time.Millisecond) // Sleep to prevent busy-waiting
		}
	}
}

func main() {
	const numOfWorkers = 4
	taskQueues := make([]chan Task, numOfWorkers)
	var wg sync.WaitGroup

	for i := range taskQueues {
		taskQueues[i] = make(chan Task, 10)
	}

	workers := make([]Worker, numOfWorkers)

	for i := range workers {
		workers[i] = Worker{
			id:       i,
			taskCh:   taskQueues[i],
			taskWg:   &wg,
			stealChs: append(taskQueues[:i], taskQueues[i+1:]...),
		}
		go workers[i].run()
	}

	for i := 0; i < 30; i++ {
		wg.Add(1)
		workerIndex := rand.Intn(numOfWorkers)
		taskQueues[workerIndex] <- func() {
			fmt.Printf("Task processed by worker %d\n", workerIndex)
		}
	}
	wg.Wait()
}
