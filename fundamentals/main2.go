package main

import (
	"fmt"
	"math"
	"runtime"
)

func sqrt(x float64) string {
	// if statement
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	//If with a short statement
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {

	sum := 0
	// for loop
	for i := 0; i < 10; i++ {
		sum += i
	}
	// init and post statements are optional
	for ; sum < 1000; {
		sum += sum
	}
	// while loop
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sqrt(2), sqrt(-4))

	// switch statement
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
	// infinite loop
	for {
	}

	//fmt.Println(sum)

}
