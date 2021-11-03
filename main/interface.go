package main

import "LearningGo"
// interface


func main() {
	r := LearningGo.Rect{Width: 10,Height: 8}
	c := LearningGo.Circle{Radius: 5}

	LearningGo.Measure(r)
	LearningGo.Measure(c)
}