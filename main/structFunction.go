package main

import (
	"fmt"
	"LearningGo"
)

//func
//func *

//func Distance(p,q LearningGo.Point) float64{
//	return math.Hypot(p.X-q.X,p.Y-q.Y)
//}

func main(){
	p := LearningGo.Point{X: 1, Y: 3}
	q := LearningGo.Point{X: 2, Y: 2}
	fmt.Println(p.Distance(q))
	p.ScaleBy(10)
	fmt.Println(p)
}
