package main

import "fmt"

// Specify new type/container
type Employee struct {
	name string
	salary int
	currency string
}

// Create methods on the Employee type

// Display salary of the emplyee
func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

// Display name of the employee
func (e Employee) displayName() {
	fmt.Printf("The name of the emplyee is %s ", e.name)
}

func main() {
	john := Employee {
		name: "John Smith",
		salary: 12000,
		currency: "$",
	}

	jane := Employee {
		name: "Jane Smith",
		salary: 15000,
		currency: "$",
	}

	john.displayName()
	fmt.Println()
	john.displaySalary()
	fmt.Println()

	jane.displayName()
	fmt.Println()
	jane.displaySalary()
	fmt.Println()
}