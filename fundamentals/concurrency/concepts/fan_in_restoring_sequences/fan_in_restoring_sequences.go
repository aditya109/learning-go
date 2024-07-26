package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) <-chan Message { // returns receive-only channel of strings
	c := make(chan Message)
	go func() {
		for i := 0; ; i++ {
			waitForIt := make(chan bool)
			c <- Message{str: fmt.Sprintf("%s %d", msg, i), wait: waitForIt}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<-waitForIt
		}
	}()

	return c
}

type Message struct {
	str  string
	wait chan bool
}

func fanin(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	go func() {
		for {
			select {
			case msg := <-input1:
				c <- msg

			case msg := <-input2:
				c <- msg
			}
		}
	}()
	return c
}

func main() {
	c := fanin(boring("joe"), boring("ann"))
	for i := 0; i < 10; i++ {
		msg := <-c
		fmt.Println(msg.str)
		msg.wait <- true // signal that the message has been processed
	}
	fmt.Println("I am done")
}
