package main

import "fmt"

type Human interface {
	Display()
}

type Contractor struct {
	name        string
	weeklyHours int
}

func (cont Contractor) Display() {
	fmt.Println("Name - ", cont.name, ", Weekly Hours - ", cont.weeklyHours)
}
