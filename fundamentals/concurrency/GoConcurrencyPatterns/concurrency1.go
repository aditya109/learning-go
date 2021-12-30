package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		var t time.Duration = time.Duration(rand.Intn(1e3)) * time.Millisecond
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(t)
	}
}

// func main() {
// 	go boring("boring!")
// }

// func main() {
//     go boring("boring!")
//     fmt.Println("I'm listening.")
//     time.Sleep(2 * time.Second)
//     fmt.Println("You're boring; I'm leaving.")
// }

func main() {
    c := make(chan string)
    go boring("boring!", c)
    for i := 0; i < 5; i++ {
        fmt.Printf("You say: %q\n", <-c) // Receive expression is just a value.
    }
    fmt.Println("You're boring; I'm leaving.")
}