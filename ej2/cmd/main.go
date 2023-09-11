package main

import (
	"app/internal"
)

func main() {
	person1 := internal.Person{
		ID:          1,
		Name:        "Cristian",
		DateOfBirth: "26/01",
	}

	employ := internal.Employee{
		ID:       2,
		Position: "Software Developer",
		Person:   person1,
	}

	employ.PrintEmployee()
}
