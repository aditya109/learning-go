package main

import (
	"fmt"
	"time"
)

/**
Design a system where a manager goroutine distributes work to multiple worker goroutines.
Each worker must send a heartbeat message back to the manager every few seconds to indicate it is still alive.
If a worker does not send a heartbeat within a specific time frame, the manager should consider the worker failed and redistribute its work to other workers.
*/

type Task struct {
	id int
}

func worker(id int, tasks <-chan Task, heartbeat chan<- int, done chan<- int) {
	for task := range tasks {
		fmt.Printf("worker %d processing task %d\n", id, task.id)
		time.Sleep(time.Second * 2)
		heartbeat <- id
		done <- task.id
	}
}

func manager(taskQueue []Task, numWorkers int) {
	tasks, heartbeat, done := make(chan Task), make(chan int), make(chan int)

	for i := 1; i <= numWorkers; i++ {
		go worker(i, tasks, heartbeat, done)
	}

	go func() {
		for _, task := range taskQueue {
			tasks <- task
		}
		close(tasks)
	}()

	workerLastSeen := make(map[int]time.Time)
	heartbeatTimeout := 3 * time.Second

	for {
		select {
		case id := <-heartbeat:
			workerLastSeen[id] = time.Now()
			fmt.Printf("Received heartbeat from worker %d\n", id)
		case taskId := <-done:
			fmt.Printf("Task %d completed\n", taskId)
		case <-time.After(heartbeatTimeout):
			now := time.Now()
			for id, lastSeen := range workerLastSeen {
				if now.Sub(lastSeen) > heartbeatTimeout {
					fmt.Printf("worker %d missed heartbeat, redistributing work\n", id)
					delete(workerLastSeen, id)
					// reassign tasks here
				}
			}
		}
		if len(workerLastSeen) == 0 {
			fmt.Println("All tasks are completed or redistributed.")
			break
		}
	}
}

func main() {
	tasks := []Task{
		{
			id: 1,
		},
		{
			id: 2,
		},
		{
			id: 3,
		},
		{
			id: 4,
		},
		{
			id: 5,
		},
		{
			id: 6,
		},
		{
			id: 7,
		},
		{
			id: 8,
		},
	}
	manager(tasks, 2)
}
