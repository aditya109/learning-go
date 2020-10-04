package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)
	go func(ch <- chan int) {
		i := <-ch // Receiving Channel data
		fmt.Println(i)
		i = <-ch // Receiving Channel data
		fmt.Println(i)
		wg.Done()
	}(ch)
	go func(ch chan <- int) {
		ch <- 42 // Sending Channel data
		ch <- 27 // error cuz #receive only channel !== #send only channels
		wg.Done()
	}(ch)
	wg.Wait()
}
