// main.go

package main

import (
	"fmt"

	"github.com/aditya109/learning-go/design-patterns/solid-patterns/1-s/example-1/problem/journal"
)

func main() {
	j := journal.Journal{}
	j.AddEntry("This is the first entry")
	j.AddEntry("This is the second entry")
	j.AddEntry("This is the third entry")
	j.AddEntry("This is the fourth entry")
	j.AddEntry("This is the fifth entry")
	fmt.Println(j)
}
