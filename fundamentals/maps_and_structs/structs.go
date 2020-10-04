package main

import "fmt"

type Doctor struct {
	Number     int
	ActorName  string
	Episodes   []string
	Companions []string
}

func main() {
	aDoctor := Doctor{
		Number:    3,
		ActorName: "Jon Pertwee",
		Companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah",
			"Jane Smith",
		},
		Episodes: []string{},
	}
	// not good syntax
	bDoctor := Doctor{
		3,
		"Jon Pertwee",
		[]string{},
		[]string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah",
			"Jane Smith",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(bDoctor)
}
