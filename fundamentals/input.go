package main

import "fmt"

func main() {
	var input string
	fmt.Println("Enter your name : ")
	fmt.Scanln(&input)
	fmt.Printf("Hello, %s!", input)
}
