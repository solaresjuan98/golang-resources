package main

import (
	"fmt"

	"github.com/donvito/hellomod"
	hellomod2 "github.com/donvito/hellomod/v2"
)

func main() {

	hellomod.SayHello()
	fmt.Println("")
	hellomod2.SayHello(" juan ")

}
