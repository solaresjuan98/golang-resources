package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full time employee"

}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

func (tEmployee TemporaryEmployee) getMessage() string {
	return "Temporary employee"
}

type PrintInfo interface {
	getMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func GetMessage(p Person) {
	fmt.Printf("%s with age %d \n", p.name, p.age)
}

func main() {

	ftEmployee := FullTimeEmployee{}

	ftEmployee.name = "Luis"
	ftEmployee.age = 2
	ftEmployee.id = 5

	fmt.Printf("%v\n", ftEmployee)

	tEmployee := TemporaryEmployee{}
	getMessage(tEmployee)
	getMessage(ftEmployee)
	//GetMessage(ftEmployee)
}
