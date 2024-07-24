package main

import (
	"fmt"
)

// You would be provided with long water string "HHOHHHOHHHOOHOHOHOHOHOO"
// you have two go routines oxygen and hydrogen, which return true, if they get "O" or "HH" respectively
// when you get "O" and "HH", you can count it as 1 water.
// Goal is to count maximum water molecules.

// splitMoleculeIntoAtom splits the molecule string into a channel of atoms.
func splitMoleculeIntoAtom(molecule string) <-chan string {
	atom := make(chan string)
	go func() {
		defer close(atom)
		for _, r := range molecule {
			atom <- string(r)
		}
	}()
	return atom
}

// tee duplicates the input channel into two output channels.
func tee(input <-chan string) (<-chan string, <-chan string) {
	out1 := make(chan string)
	out2 := make(chan string)
	go func() {
		defer close(out1)
		defer close(out2)
		for item := range input {
			// Send the item to both channels
			select {
			case out1 <- item:
			case out2 <- item:
			}
			select {
			case out1 <- item:
			case out2 <- item:
			}
		}
	}()
	return out1, out2
}

// oxygen processes the atoms to count oxygen atoms.
func oxygen(atom <-chan string) <-chan bool {
	oxChan := make(chan bool)
	go func() {
		defer close(oxChan)
		for a := range atom {
			if a == "O" {
				oxChan <- true
			}
		}
	}()
	return oxChan
}

// hydrogen processes the atoms to count hydrogen pairs.
func hydrogen(atom <-chan string) <-chan bool {
	hyChan := make(chan bool)
	go func() {
		defer close(hyChan)
		hCount := 0
		for a := range atom {
			if a == "H" {
				hCount++
				if hCount == 2 {
					hyChan <- true
					hCount = 0
				}
			}
		}
	}()
	return hyChan
}

// counter counts the number of water molecules that can be formed.
func counter(oxChan, hyChan <-chan bool) int {
	var waterCount = 0
	done := make(chan bool)
	go func() {
		for {
			select {
			case _, ok := <-oxChan:
				if !ok {
					done <- true
					return
				}
				select {
				case _, ok := <-hyChan:
					if !ok {
						done <- true
						return
					}
					waterCount++
				}
			}
		}
	}()
	<-done
	return waterCount
}

func main() {
	input := "HHOHHHOHHHOOHOHOHOHOHOO"

	atom := splitMoleculeIntoAtom(input)
	oxChan, hyChan := tee(atom)

	oxProcessed := oxygen(oxChan)
	hyProcessed := hydrogen(hyChan)

	waterCount := counter(oxProcessed, hyProcessed)

	fmt.Printf("Number of water molecules: %d\n", waterCount)
}
