package main

import(
	"fmt"
)

func foo(x *[3]int) int{
	(*x)[0] = (*x)[0] + 1
	return (*x)[0]
}

func foo2(sli []int) int{
	sli[2] = sli[2] + 1
	return sli[0]
}
func main(){
	// passando um array para poder ser alterado na função
	a := [3]int{1, 2, 3}
	foo(&a)
	fmt.Print(a)
	fmt.Printf(": passando arrays como arrays.\n\n")
	// out : [2, 2, 3]

	// no geral o método de passar os ponteiros do array é o slice, da seguinte maneira:
	b := []int{1, 2, 3}
	foo2(b)
	fmt.Print(b)
	fmt.Printf(": passando arrays como slices.\n\n")
	// out : [1, 2, 4]
}

