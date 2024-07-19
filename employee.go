package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	age       int
}

func (emp Employee) Display() {
	fmt.Printf("Hello from Employee First Name: %s Last Name: %s Age %d \n", emp.FirstName, emp.LastName, emp.age)
}

// Methods for struct Employee – this is pointer receiver type
func (emp *Employee) increaseAgeByOne() {
	emp.age++
}
func (emp *Employee) increaseAge(increaseBy int) {
	emp.age += increaseBy
}
