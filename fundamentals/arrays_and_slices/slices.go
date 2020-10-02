package main

import "fmt"

func main() {
	a:= []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//fmt.Println(a)
	//fmt.Println(len(a))
	//fmt.Printf("Capacity: %v\n", cap(a))

	b := a[:]
	c := a[3:]
	d := a[:6]
	e := a[3:6]

	fmt.Println(a, b, c, d, e)
	a[5] = 43

	fmt.Println(a, b, c, d, e)
}
