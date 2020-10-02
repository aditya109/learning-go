package display

import "fmt"

//DisplayArrayOrSlice is used to display array or slice
func DisplayArrayOrSlice(a []int, name string) {
	fmt.Printf("Array `%v` contains : %v \n", name, a)
	fmt.Printf("Length : %v\n", len(a))
	fmt.Printf("Capacity: %v\n\n", cap(a))
}