package main

import (
	"fmt"
)

func main() {

	slice := []string{"apple", "banana", "orange", "date"}
	result := convertToMap(slice)
	fmt.Println(result)

}

func convertToMap(items []string) map[string]float64 {

	// Create a map object
	result := make(map[string]float64)
	// Your code goes here
	return result
}
