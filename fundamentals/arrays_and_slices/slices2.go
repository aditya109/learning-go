package main

import (
	"github.com/aditya109/fundamentals/display"
)

func main() {
	a := make([]int, 3)
	display.DisplayArrayOrSlice(a, "a")

	b := make([]int, 3, 100)
	display.DisplayArrayOrSlice(b, "b")

	c := []int{} // or var c []int
	display.DisplayArrayOrSlice(c, "c")

	c = append(a, 1)
	display.DisplayArrayOrSlice(c, "c")

	c = append(a, 2, 3, 4, 5, 6)
	display.DisplayArrayOrSlice(c, "c")

	c = append(c, b...)
	display.DisplayArrayOrSlice(c, "c")
}
