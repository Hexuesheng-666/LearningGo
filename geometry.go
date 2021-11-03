package LearningGo

import (
	"fmt"
	"math"
)

type Geometry interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	Width,Height float64
}

type Circle struct {
	Radius float64
}

func (r Rect) Area() float64{
	return r.Width * r.Height
}

func (r Rect) Perimeter() float64{
	return 2 * ( r.Width + r.Height)
}

func (c Circle) Area() float64{
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64{
	return 2 * math.Pi * c.Radius
}

func Measure(g Geometry){
	fmt.Println(g)
	fmt.Println("Area = " , g.Area())
	fmt.Println("Perimeter = " , g.Perimeter())
}