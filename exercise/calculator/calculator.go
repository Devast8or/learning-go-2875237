package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter first value: ")
	value1, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Print("Enter second value: ")
	value2, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	calculate(value1, value2)
}

// calculate() returns the sum of the two parameters
func calculate(value1 string, value2 string) float64 {

	aFloat, err := strconv.ParseFloat(strings.TrimSpace(value1), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value must be a number")
	} else {
		fmt.Println("Value of number:", aFloat)
	}

	aFloat2, err := strconv.ParseFloat(strings.TrimSpace(value2), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value must be a number")
	} else {
		fmt.Println("Value of number:", aFloat2)
	}

	result := aFloat + aFloat2

	return result
}
