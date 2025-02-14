package main

import "fmt"

type Person struct{
	name string
	addr string
	phone string
}

func main(){

	x1 := new(Person)
	var x2 int

	// p n dar erro dps
	_ = x2
	
	
	x1.name = "joe"
	fmt.Println(x1.name)

	// também pode ser inicializado da seguinte maneira
	// se atentar que os valores estão entre chaves !!!!
	x3 := Person{name: "Bruno", addr:"bleble", phone:"12345"}
	fmt.Println(x3)
}