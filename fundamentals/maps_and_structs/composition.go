package main

import "fmt"

type Animal struct {
	Name 	string
	Origin 	string
}

type Bird struct {
	Animal
	SpeedKPH 	float32
	CanFly		bool
}

func main()  {
	b  := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)

	b2 := Bird {
		Animal: Animal{
			Name: "Emu",
			Origin: "Australia",
		},
		SpeedKPH: 30,
		CanFly: false,
	}
	fmt.Println(b2)
}