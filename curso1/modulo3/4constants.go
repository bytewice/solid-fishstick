package main

import "fmt"

// expressões cujos valores serão conhecidos durante a compilaçao
// o tipo é inferido do que estiver do lado direito como se tivesse um :=
const x = 1.3
const (
	y = 4
	z = "Hi"
)

// iota é como se fosse enumerate
type Grades int
const(
	A Grades = iota
	B
	C
	D
	E
	F
)

func main(){
	// ...
	// ...
	fmt.Printf("%d %d %d %d %d %d\n", A, B, C, D, E, F)
	// prints 0 1 2 3 4 
}