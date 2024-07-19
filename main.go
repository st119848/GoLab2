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
	emp.Display()
	var cont Human = Contractor{name: "XYZ", weeklyHours: 35}
	cont.Display()

}
