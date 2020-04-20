package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"time"
)

const (
	layoutISO = "2006-01-02"
)

// Person struct describes every from array
type Person struct {
	firstName string
	lastName  string
	birthDay  time.Time
}

// People - slice of people
type People []Person

func (p People) Len() int {
	return len(p)
}

func (p People) Less(i, j int) bool {
	if p[i].birthDay.Equal(p[j].birthDay) {
		if p[i].firstName == p[j].firstName {
			return p[i].lastName < p[j].lastName
		}
		return p[i].firstName < p[j].firstName
	}
	return p[i].birthDay.Before(p[j].birthDay)
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func makeDateTime(s string) (birthDate time.Time) {
	birthDate, err := time.Parse(layoutISO, s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return
}

// Figure - interface for all graphic structures
type Figure interface {
	area() float64
	perimeter() float64
}

// Square is equilateral rectangle with side
type Square struct {
	side int
}

// Circle is sample figure with radius
type Circle struct {
	radius int
}

func (s Square) area() float64 {
	return math.Pow(float64(s.side), 2)
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(float64(c.radius), 2)
}

func (s Square) perimeter() float64 {
	return 4 * float64(s.side)
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * float64(c.radius)
}

func main() {
	// Task 1
	p := People{
		{"Ivan", "Ivanov", makeDateTime("2003-12-12")},
		{"Ivan", "Ivanov", makeDateTime("2004-07-02")},
		{"Artiom", "Ivanov", makeDateTime("2003-12-12")},
		{"Bob", "Ivanov", makeDateTime("2003-12-12")},
		{"Bob", "Ivaniw", makeDateTime("2003-12-12")},
	}

	sort.Sort(p)
	fmt.Println(p)

	// Task 2
	var s Figure = Square{side: 2}
	var c Figure = Circle{radius: 3}

	fmt.Println(s.area(), s.perimeter())
	fmt.Println(c.area(), c.perimeter())
}
