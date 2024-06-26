package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string, quit chan bool) <-chan string { // returns receive-only channel of strings
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s: %d", msg, i):
				time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			case <-quit:
				return
			}
		}
	}()

	return c
}

func main() {
	quit := make(chan bool)
	c := boring("Hello World", quit)
	for i := rand.Intn(10); i > 0; i-- {
		fmt.Println(<-c)
	}
	quit <- true
	fmt.Println("i am done now")
}
