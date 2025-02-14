package main

import "fmt"

// if e while funciona igual a C

// Break sai do loop 

// continue skipa so a iteração do loop

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
		fmt.Print(value + y)
		// fmt.Print(value + y + z)
		// não funciona pq junta tipos de dados diferentes
		fmt.Print(z)
	}
	}
}

func switch_non_tag(){

	switch:{
	case x<1:{
		fmt.Printf("Case 1")
		}
	case x>1:{
		fmt.Printf("Case 2")
		}
	default:{
		fmt.Printf("Default case")
		}
	}
}

func scanear(){
	var num_apple int

	fmt.Printf("What's the number of apples?")
	scan(&num_apple)
	// n sei a diferença entre scan e scanln
	// aparentemente scanln é scan line, mas lê até encontrar o primeiro 'espaço'

	fmt.Printf(num_apple)
}