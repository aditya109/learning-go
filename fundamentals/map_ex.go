package main

import (
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	sp := strings.Split(s, " ")
	for i := range sp {
		if _, found := m[sp[i]]; found {
			m[sp[i]] += 1
		} else {
			m[sp[i]] = 1
		}
	}
	return m
}

func main() {
	WordCount("a boy a girl kicked boy go girl")
	//wc.Test(WordCount)
}
