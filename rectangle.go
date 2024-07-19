package main

import "fmt"

type Rectangle struct {
	Height float32
	Width  float32
}

// Method that returns a result
func (rect Rectangle) Area() float32 {
	return rect.Height * rect.Width
}

func (rect *Rectangle) increaseWidth(w float32) {
	rect.Width = rect.Width + w
}

func (rect Rectangle) Display() {
	fmt.Print("Width: ", rect.Width, " Height: ", rect.Height)
}
