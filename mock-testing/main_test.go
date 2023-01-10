package main

import "testing"

func TestGetFullTimeEmployeeById(t *testing.T) {

	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "1",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						Id:       1,
						Position: "CEO",
					}, nil
				}

				GetPersonByDNI = func(id string) (Person, error) {
					return Person{
						Name: "John Doe",
						Age:  35,
						DNI:  "1",
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					Age:  35,
					DNI:  "1",
					Name: "John Doe",
				},
				Employee: Employee{
					Id:       1,
					Position: "CEO",
				},
			},
		},
	}
	origingalGetEmployeeById := GetEmployeeById
	origingalGetPersonByDni := GetPersonByDNI

	for _, test := range table {
		test.mockFunc()

		ft, err := GetFullTimeEmployeeById(test.id, test.dni)

		if err != nil {
			t.Errorf("Error when getting employee")
		}

		if ft.Age != test.expectedEmployee.Age {
			t.Errorf("Error, got %d exepected %d", ft.Age, test.expectedEmployee.Age)
		}

	}

	GetEmployeeById = origingalGetEmployeeById
	GetPersonByDNI = origingalGetPersonByDni

}
