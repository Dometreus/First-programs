package main

import "fmt"

var ppl int
var ok bool

func tORf() { // cheking for state

	if ok == true {

		fmt.Println(":D")
		fmt.Println("the information you entered is", ok)
		fmt.Println("there are", ppl, "people in that area")

	} else {

		fmt.Println(":O")
		fmt.Println("The information you entered is", ok)
		fmt.Println("There's no info for that state in this database")
		fmt.Println("Please check for errors or enter information")

	}

}

func main() {
	statePopulations := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
	}

	statePopulations["New York"] = 49745289 // add states here

	//delete(statePopulations, "New York")  delete stetes here

	fmt.Println("the database has", len(statePopulations), "states") // lenghth of map

	ppl, ok = statePopulations["California"] //check individual state info

	tORf()

}
