package main

import (
	"fmt"
	"github.com/aditya109/fundamentals/test1"
)

func add(x int, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func absoluteDistance(x, y int) (int, int) {
	return x - y, x + y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool

func main() {
	k := 34
	var i = 3
	c, python, java := true, false, "no!"

	fmt.Println(add(42, 13))
	fmt.Println(sub(42, 13))
	a, b := absoluteDistance(42, 13)
	fmt.Println(a, b)
	fmt.Println(split(17))
	fmt.Println(i, k, c, java, python)
	fmt.Println("Who said Tea ? ", test1.WhoSaidTea)
}
