package main

import(
	"encapsulation/utils"
)

func main(){
	var p data.Point
	p.InitMe(3,4)
	p.Scale(3)
	p.PrintMe()
}