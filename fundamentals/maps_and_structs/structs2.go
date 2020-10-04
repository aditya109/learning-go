package main

import "fmt"

func main() {
	// Anonymous Struct
	aDoctor := struct {name string} {name: "John Pertwee"}
	anotherDoctor := &aDoctor
	anotherDoctor.name = "Tom Baker"
	fmt.Println(anotherDoctor)
	fmt.Println(aDoctor)
}
