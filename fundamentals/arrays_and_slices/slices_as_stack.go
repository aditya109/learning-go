package main

import (
	"github.com/aditya109/fundamentals/display"
)

func main() {

	a:= []int{1, 2, 3, 4, 5}
	display.DisplayArrayOrSlice(a, "a")

	b:= a[1:len(a)]
	display.DisplayArrayOrSlice(b, "b")
	display.DisplayArrayOrSlice(a, "a")
	// IMPORTANT !
	c:= append(a[:2], a[3:]...)
	display.DisplayArrayOrSlice(c, "c")
	display.DisplayArrayOrSlice(a, "a")
}
