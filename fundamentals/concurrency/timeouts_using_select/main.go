package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

// block channel for specified duration of time

//func main() {
//	c := boring("hello world")
//
//	for {
//		select {
//		case msg := <-c:
//			fmt.Println(msg)
//		case <-time.After(500 * time.Millisecond):
//			fmt.Println("you are too slow")
//			return
//		}
//	}
//}

// Block the entire conversation for specified amount of time

func main() {
	c := boring("Hello World")
	timeout := time.After(3 * time.Second)
	for {
		select {
		case msg := <-c:
			fmt.Println(msg)
		case <-timeout:
			fmt.Println("you all are too slow")
			return
		}
	}
}
