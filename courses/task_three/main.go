package main

import "fmt"

// Point type represents a point in Cartesian coordinate system
type Point struct {
	x, y int
}

// Square type is equilateral rectangle
type Square struct {
	start Point // start position in Cartesian coordinate system
	a     uint  // side of a square
}

// Perimeter method - calculate perimetr for Square
func (s Square) Perimeter() int {
	return int(s.a) * 4
}

// Area method - calculate area for Square
func (s Square) Area() int {
	return int(s.a) * int(s.a)
}

// End method - return Point type end position in Cartesian coordinate system
func (s Square) End() Point {
	return Point{
		x: s.start.x + int(s.a),
		y: s.start.y - int(s.a),
	}
}

func main() {
	s := Square{Point{1, 1}, 5}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())
}
