package main

import (
	"fmt"
)

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

func UseIt(size Sized) {
	width := size.GetWidth()
	size.SetHeight(10)
	expected_area := 10 * width
	actual_area := size.GetHeight() * size.GetWidth()
	fmt.Print("Expected and are of ", expected_area, " but got ", actual_area, "\n")
}

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}

func main() {
	rc := &Rectangle{2, 3}
	UseIt(rc)
	sq := NewSquare(5)
	UseIt(sq)
}
