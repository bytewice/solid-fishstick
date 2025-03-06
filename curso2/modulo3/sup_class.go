package main

import(
	"fmt"
	"math"
)

type Point struct{
	x float64
	y float64
}

func (p Point) DistToOrig() int{
	
	t:= math.Pow(p.x, 2) + math.Pow(p.y, 2)
	return int(math.Sqrt(t))
}

func main(){

	p1 := Point{8, 6}
	fmt.Println(p1.x, p1.y)
	fmt.Println(p1.DistToOrig())
}