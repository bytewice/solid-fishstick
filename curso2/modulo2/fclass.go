package main

import(
	"fmt"
)

// utilizando funções como variáveis

// função 
var funcVar func(int) int

func incFn(x int) int{
	return x+1
}

func applyIt( afucnt func (int) int, val int) int{
	return afucnt(val)
}

func decFn (x int) int { return x-1}


func main(){
	// basicamente funcVar virou incFn momentaneamente (( n sei qual a vantagem disso, parece doidera))
	funcVar = incFn
	fmt.Println(funcVar(1))
	fmt.Println(incFn(1))

	// exemplo de como uma função pode ser argumento de outra função
	fmt.Println(applyIt(incFn, 2))
	fmt.Println(applyIt(decFn , 2))

	// na hora de passar a função como argumento, o argumento pode ate ser a declaração de uma função:
	v := applyIt(func (x int) int {return x+1}, 11)
	fmt.Println(v)
}