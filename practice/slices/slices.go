package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)
	colors = append(colors, "Purple")
	fmt.Println(colors)

	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	numbers := make([]int, 5, 5)
	numbers[0] = 134
	numbers[1] = 1
	numbers[2] = 34
	numbers[3] = 4
	numbers[4] = 3
	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)

}
