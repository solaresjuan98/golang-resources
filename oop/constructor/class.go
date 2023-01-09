package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {

	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}

}

func main() {

	// * 1
	e := Employee{}
	fmt.Printf("%v\n", e)

	// * 2
	e2 := Employee{
		id:       1,
		name:     "Juan",
		vacation: true,
	}

	fmt.Printf("%v\n", e2)

	// * 3
	e3 := new(Employee)

	fmt.Printf("%v\n", *e3)

	e3.id = 3
	e3.name = "Antonio"
	e3.vacation = false

	// * 4
	e4 := NewEmployee(25, "Pablo", false)

	fmt.Printf("%v\n", *e4)

}
