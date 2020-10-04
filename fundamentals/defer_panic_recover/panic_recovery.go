package main

import (
	"fmt"
	"log"
)

func panicked() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("error : ", err)
			panic(err)
		}
	}()
	panic("something bad happened  ")
	fmt.Println("done panicking")
}

func main() {
	fmt.Println("start")
	panicked()
	fmt.Println("end")
}
