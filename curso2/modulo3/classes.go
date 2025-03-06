package main

import(
	"fmt"
)	

//  class Point:
//  	def _init_ (self, xval, yval):
//  		self.x = xval
//  		self.y = yval	 

// Associado métodos com os dados
type MyInt int

func (mi MyInt) Double () int{
	return int(mi*2)
}

func main(){
	v:= MyInt(3)
	fmt.Println(v.Double())
}