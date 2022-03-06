// example-1/solution/journal/journal.go

package journal

import (
	"fmt"
)

type Journal struct {
	entries []string
	count   int
}

func (J *Journal) AddEntry(entry string) {
	J.count++
	J.entries = append(J.entries, entry)
}

func (J *Journal) DeleteEntry(position int) {
	J.entries = append(J.entries[:position], J.entries[position+1:]...)
}

func (J Journal) String() string {
	var result string
	for i := 0; i < J.count; i++ {
		result += fmt.Sprintf("Entry %d: %s", i+1, J.entries[i])
		result += fmt.Sprint("\n")
	}
	return result
}

