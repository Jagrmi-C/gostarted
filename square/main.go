package main

import "fmt"

// Point is position in space
type Point struct {
	x, y int
}

// Square is equilateral rectangle
type Square struct {
	start Point
	a     uint
}

// Perimeter - calculate perimetr for Square
func (s Square) Perimeter() int {
	return int(s.a) * 4
}

// Area - calculate area for Square
func (s Square) Area() int {
	return int(s.a) * int(s.a)
}

// End - Point after made ...
func (s Square) End() Point {
	newX := s.start.x + int(s.a)
	newY := s.start.y - int(s.a)
	endPoint := Point{newX, newY}
	return endPoint
}

func main() {
	s := Square{Point{1, 1}, 5}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())
}
