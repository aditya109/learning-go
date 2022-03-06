// example-1/problem/journal/journal.go

package journal

import (
	"fmt"
	"os"
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

// Adding these new functions
func (J Journal) SaveToFile(filename string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, 0666)
	check(err)
	defer file.Close()
	_, err = file.WriteString(J.String())
	check(err)
}

func (J Journal) LoadFileFromDisk(filename string) {
	// implementation pending
}

func (J Journal) LoadFileFromWeb(url string) {
	// implementation pending
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
