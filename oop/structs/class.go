package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func (e *Employee) SetId(id int) {
	e.id = id

}

func (e *Employee) SetName(name string) {
	e.name = name

}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func main() {
	e := Employee{}

	e.name = "Juan"
	e.SetId(2)
	//fmt.Printf("%v", e)

	fmt.Println(e.GetId())
	fmt.Println(e.GetName())

}
