package main

import (
	"fmt"
	"time"
)

// create a "ping/pong" game with "buffered" channels.
// There will be two goroutines which are dedicated to ping and pong the ball every second infinitely.
// Channels are used for communication among goroutines so in this case two goroutines will inform other once they hit the ball.

func Ping(ping chan bool, pong chan bool) {
	for {
		select {
		case <-ping:
			fmt.Println("Ping")
			time.Sleep(1000 * time.Millisecond)
			pong <- true
		}
	}
}

func Pong(ping chan bool, pong chan bool) {
	for {
		select {
		case <-pong:
			fmt.Println("Pong")
			time.Sleep(1000 * time.Millisecond)
			ping <- true
		}
	}
}

func main() {
	ping, pong := make(chan bool), make(chan bool)

	go Ping(ping, pong)
	go Pong(ping, pong)

	ping <- true

	time.Sleep(1000 * time.Second)

}
