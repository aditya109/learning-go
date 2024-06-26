package main

import (
	"fmt"
	"sync"
)

type SafeMapConstruct struct {
	mu     sync.Mutex
	numMap map[string]int
}

func (s *SafeMapConstruct) Add(num int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.numMap["key"] = num
}

func main() {
	s := &SafeMapConstruct{
		numMap: make(map[string]int),
	}

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.Add(i)
		}(i)
	}
	wg.Wait()
	fmt.Println(s.numMap["key"])
}
