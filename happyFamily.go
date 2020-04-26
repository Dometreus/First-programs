package main

import "fmt"

type family struct {
	name     string
	age      int
	children []string
}

func main() {

	fmt.Println("........................................................................................")
	fmt.Println(" This is a program that tells u")
	fmt.Println("information about the family members")
	fmt.Println("|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||")

	fmt.Println(" (name) (age)     (children info)")

	fMem := family{
		name: "Father",
		age:  45,
		children: []string{
			"Generic boy name here",
			"Generic girl name here",
		},
	}

	mMem := family{
		name: "Mother",
		age:  38,
		children: []string{
			"Generic boy name here",
			"Generic girl name here",
		},
	}

	bMem := family{
		name: "little Peter",
		age:  3,
		children: []string{
			"cant make them yet",
		},
	}

	gMem := family{
		name: "Emma",
		age:  17,
		children: []string{
			"luckily none",
		},
	}

	fmt.Println(fMem)
	fmt.Println(mMem)
	fmt.Println(bMem)
	fmt.Println(gMem)

}
