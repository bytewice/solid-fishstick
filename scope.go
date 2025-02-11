package main

import "fmt"

var x = 4
//funciona parecido com c também

func f(){
	//iniciar em uma função não dá autoridade nas outras funções se eu não passar pra elas
	y := 3
	fmt.Printf("%d", x)
}

func g(){
	//g não tem autoridade para printar y
	// fmt.Printf("")
	fmt.Printf("%d", x)
}


func g{

	x := 4
	fmt.Printf("%d", x)

	//nesse caso func g consegue printar x pois o bloco de func p está contido no bloco de func g
	func p(){
		fmt.Printf("%d", x)
	}
}