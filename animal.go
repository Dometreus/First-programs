package main

import "fmt"

type animal struct {
	name   string
	height float32
}

type bird struct {
	animal
	fliyingSpeed float32
	offspring    int
}

type mammal struct {
	animal
	runningSpeed float32
	offspring    int
}

func main() {

	individual1 := bird{
		animal:       animal{name: "hummingbird", height: 0.05},
		fliyingSpeed: 20,
		offspring:    12,
	}

	individual2 := mammal{
		animal:       animal{name: "dog", height: 0.5},
		runningSpeed: 35,
		offspring:    0, //castrated
	}

	fmt.Println(individual1)
	fmt.Println(individual2)
}
