package main

import "fmt"

// "<-"" in parameters means the channel is only used to write
func Generator(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}

	close(c)

}

// "<-"" in parameters means the channel is only used to read
func Double(in <-chan int, out chan int) {

	for value := range in {
		out <- 2 * value
	}

	close(out)

}

func Print(c <-chan int) {
	for value := range c {
		fmt.Println(value)
	}
}

func main() {
	generator := make(chan int)
	doubles := make(chan int)

	go Generator(generator)
	go Double(generator, doubles)

	Print(doubles)

}
