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

//func main() {
//	joe := boring("joe")
//	ann := boring("ann")
//
//	for i := 0; i < 10; i++ {
//		fmt.Println(<-joe, <-ann)
//	}
//}

func fanin(input1 <-chan string, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func main() {
	c := fanin(boring("joe"), boring("ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("i am done")
}
