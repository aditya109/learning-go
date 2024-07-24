package main

import (
	"fmt"
	"sync"
	"time"
)

// create a "ping/pong" game with "buffered" channels.
// There will be two goroutines which are dedicated to ping and pong the ball every second infinitely.
// Channels are used for communication among goroutines so in this case two goroutines will inform other once they hit the ball.

func main() {
	ping, pong := make(chan bool), make(chan bool)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go Ping(ping, pong, wg)
	go Pong(ping, pong, wg)

	ping <- true

	wg.Wait()

}

func Ping(ping, pong chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ping:
			fmt.Println("Ping")
			time.Sleep(1000 * time.Millisecond)
			pong <- true
		}
	}
}

func Pong(ping, pong chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-pong:
			fmt.Println("Pong")
			time.Sleep(1000 * time.Millisecond)
			ping <- true
		}
	}
}
