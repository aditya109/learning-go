package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("We are recovering ...")
			log.Println("Error : ", err)
		}
	}()
	panic("Something Bad Happened !")
	fmt.Println("end")
}
