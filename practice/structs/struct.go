package main

import (
	"fmt"
)

func main() {
	poodle := Dog{"Poodle", 10}
	fmt.Println("Structs")
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed %v\nWeight: %v\n", poodle.Breed, poodle.Weight)
	poodle.Weight = 9
	fmt.Printf("Breed %v\nWeight: %v\n", poodle.Breed, poodle.Weight)

}

// Dog is a struct
type Dog struct {
	Breed  string //Uppercase exported
	Weight int
}

//Custom structure (dog would be private)
