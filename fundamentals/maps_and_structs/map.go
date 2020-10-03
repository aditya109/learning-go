package main

import "fmt"

func main() {
	statePopulations := make(map[string] int)
	statePopulations = map[string]int {
		"Marlene" 	: 1,
		"Alec" 		: 2,
		"Hiram"		: 3,
		"Benedict" 	: 4,
		"Claudio" 	: 5,
	}
	statePopulations["Aryan"] = 6
	fmt.Println(statePopulations)
	// return order of map is not fixed
	delete(statePopulations, "Marlene")
	fmt.Println(statePopulations)

	fmt.Println(len(statePopulations))
	pop, ok := statePopulations["Hiram"]
	fmt.Println(pop, ok)
	pop, ok = statePopulations["Beny"]
	fmt.Println(pop, ok)
}

