package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	var x int
	x = 8
	y := 7.0

	fmt.Println(x)
	fmt.Println(y)

	// * Capture value and error
	myValue, err := strconv.ParseInt("NaN", 0, 64)

	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(myValue)
	}

	// * Map key - value
	m := make(map[string]int)
	m["Key"] = 6
	fmt.Println(m["Key"])

	fmt.Println("=============")
	// * Simple array
	s := []int{1, 2, 3}
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}

	// * Append an element to the array
	s = append(s, 16)
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}
	fmt.Println("===== Go routine =====")
	// Create a channel
	// c := make(chan int)

	// // * Go routines
	// go doSomething(c)
	// <-c

	g := 25
	fmt.Println(g)

	h := &g
	fmt.Println(h)
	fmt.Println(*h)

}

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	//
	c <- 1
}
