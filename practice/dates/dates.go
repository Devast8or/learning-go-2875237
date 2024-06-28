package main

import (
	"fmt"
	"time"
)

func main() {

	n := time.Now()
	fmt.Println("I recorded this video at ", n)

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Go Launched at ", t)
	fmt.Println(t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 23:00:00 2009")
	fmt.Printf("The type of parsed Time is %T\n", parsedTime)

}
