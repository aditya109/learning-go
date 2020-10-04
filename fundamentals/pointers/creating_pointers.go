package main

import "fmt"

func main() {
	a := []int {1,2,3}
	b:=a
	fmt.Println(a, b)
	a[1] = 42
	fmt.Println(a, b)
}

type myStruct struct {
	foo int
}
