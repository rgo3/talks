package main

import (
	"fmt"
	"math"
)

// START OMIT
// Any type with an Area method satisfies this interface.
type Arear interface {
	Area() float64
}

type Circle struct{ Radius float64 }
type Square struct {
	Hight float64
	Width float64
}

func (c Circle) Area() float64  { return math.Pi * math.Pow(c.Radius, 2) }
func (s *Square) Area() float64 { return s.Hight * s.Width }

func main() {
	c := Circle{Radius: 2}
	s := &Square{2, 4}
	printArea(c)
	printArea(s)
}

func printArea(a Arear) {
	fmt.Println("area:", a.Area())
}

// END OMIT
