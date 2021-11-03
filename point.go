package LearningGo

import "math"


type Point struct {
	X float64
	Y float64
}

func (p Point) Distance(q Point) float64{
	return math.Hypot(p.X-q.X,p.Y-q.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}