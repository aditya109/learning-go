package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func boring(msg string) { // returns receive-only channel of strings
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Printf("%s %d\n", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
}

func main() {
	wg.Add(1)
	go boring("joe!")

	wg.Add(1)
	go boring("add!")

	wg.Wait()
	fmt.Println("done")
}
