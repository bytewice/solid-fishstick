package main

import "fmt" 

func main(){
	// os operadores de ponteiros são parecidos com c: & e *
	var x int = 1
	var y int

	// ip aponta para o endereço de um int
	var ip *int 

	// esse endereço vai ser o endereço de x
	ip = &x 

	// (*ip) é o conteúdo que tem no endereço apontado por ip
	y = *ip
}

func not_main(){
	// iniciação curta de ponteiros
	// marca assinada do golang
	ptr := new(int)
	*ptr = 3
}