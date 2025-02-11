package main

import "fmt"

// if e while funciona igual a C

func main(){
	// também pode iniciar a variável no começo do for q nem em C
	for i:=3; i<10; i++{
		fmt.Printf("%d\n", i)
	}

	// Não existe while em go mas o for pode ser usado da seguinte maneira
	// que aí funciona igual o while
	i := 0
	for i < 10{
		fmt.Printf("%d ", i)
		i++
	}

	fmt.Printf("\n")

	y := 122
	x := 123
	z := "ablubleble"
	switch x{
	case 1: fmt.Printf("case 1")
	case 2: fmt.Printf("case 2")
	case 123: {
		value := 123
		fmt.Printf("Case %d\n", value)
		fmt.Print(value + y + z)
		//fmt.Print(z)
	}
	}
}