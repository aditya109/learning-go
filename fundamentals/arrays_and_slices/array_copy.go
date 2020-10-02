package main

import "fmt"

func main() {

	a:= [...]int{1, 2, 3}
	// Deep copy
	b:= a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	c:=&a
	c[1] = 5
	fmt.Println(a)
	fmt.Println(c)

}
