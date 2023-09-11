package internal

import "fmt"

type Employee struct {
	ID       int
	Position string
	Person
}

func (e Employee) PrintEmployee() {
	fmt.Println("ID Employee: ", e.ID)
	fmt.Println("Position Employee: ", e.Position)
	fmt.Println("ID Person: ", e.Person.ID)
	fmt.Println("Name Person: ", e.Person.Name)
	fmt.Println("DateOfBirth Person: ", e.Person.DateOfBirth)
}
