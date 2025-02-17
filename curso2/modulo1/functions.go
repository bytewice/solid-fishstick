package main

import (
	"fmt"
)

// tanto go quanto C utiliza um método call by value que para alterar o valor da variável
// na função, é necessário passar um ponteiro para essa variável!
func second(x *int){
	fmt.Printf("Hello world! no: %d\n", *x)
}

func main(){
	x := 4
	second(&x)
}