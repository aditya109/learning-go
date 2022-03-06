// example-1/solution/persistancemanager/persistancemanager.go

package persistancemanager

import (
	"os"

	"github.com/aditya109/learning-go/design-patterns/solid-patterns/1-s/example-1/solution/journal"
)

// Adding these new functions
func SaveToFile(filename string, journal journal.Journal) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, 0666)
	check(err)
	defer file.Close()
	_, err = file.WriteString(journal.String())
	check(err)
}

func LoadFileFromDisk(filename string, journal journal.Journal) {
	// implementation pending
}

func LoadFileFromWeb(url string, journal journal.Journal) {
	// implementation pending
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
