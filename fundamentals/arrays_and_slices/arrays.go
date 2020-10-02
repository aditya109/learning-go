package main

import "fmt"

func main() {
	grades := [3]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades)

	grades2 := [...]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades2)

	var grades3 [3]int
	fmt.Printf("Grades: %v\n", grades3)
	grades3[0] = 97
	fmt.Printf("Grades: %v\n", grades3)
	grades3[1] = 85
	fmt.Printf("Grades: %v\n", grades3)
	grades3[2] = 93
	fmt.Printf("Grades: %v\n", grades3)

	var len int = len(grades3)
	fmt.Printf("Length: %v\n", len)
}
