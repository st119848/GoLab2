package main

func main() {
	// emp := Employee{"John", "Wicker", 10}
	// emp.Display()
	// emp.increaseAgeByOne()
	// emp.Display()
	// emp.increaseAge(5)
	// emp.Display()

	// rect := Rectangle{Height: 25.5, Width: 12.75}
	// fmt.Println(rect.Area())
	// rect.increaseWidth(10.0)
	// fmt.Println(rect)

	var emp Human = Employee{name: "Rob Pike", designation: "Engineer"}
	var emp2 Human = Employee{name: "John", designation: "Scientist"}

	var cont Human = Contractor{name: "XYZ", weeklyHours: 35}
	var cont2 Human = Contractor{name: "ABC", weeklyHours: 40}

	humans := []Human{emp, emp2, cont, cont2}

	for i := 0; i < len(humans); i++ {
		humans[i].Display()
	}

	for _, human := range humans {
		human.Display()
	}

}
