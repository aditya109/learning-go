// goroutines1.go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

//func boring(msg string, c chan string) {
//	for i := 0; ; i++ {
//		c <- fmt.Sprintf("%s %d", msg, i)
//		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
//	}
//}
//
//func main() {
//	c := make(chan string)
//	go boring("boring", c)
//	for i := 0; i < 5; i++ {
//		fmt.Printf("you say: %q\n", <-c)
//	}
//	fmt.Println("I'm done")
//}

// generator: function that returns a channel

func boring(msg string) <-chan string { // returns receive-only channel of strings
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

func main() {
	c := boring("boring goroutine")
	for msg := range c {
		fmt.Println(msg)
	}
	fmt.Println("I'm done")
}
