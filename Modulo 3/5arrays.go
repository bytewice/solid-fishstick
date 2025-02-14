package main

import "fmt"

// tomar cuidado com erro de variável não usada!

func arrays() {
    // Os elementos são iniciados com valor 0!
    var x [5]int

    x[0] = 2
    fmt.Printf("%d\n", x[1]) 
	// Prints 0, pq os valores são inicializados como 0

    y := [...]int{1, 2, 3, 4}
    // Evita erro de variável não usada
	_ = y
}

func iterating() {
    x := [3]int{1, 2, 4}

    for i, v := range x {
        fmt.Printf("ind %d, Val %d\n", i, v)
    }
}

// Entendo como o range funciona por debaixo dos panos
func range_(){
	// esse for é equivalente ao range utilizado na função logo acima
	for i := 0; i < len(x); i++ {
		v := x[i] // Pegando o valor do índice atual
		fmt.Printf("Índice: %d, Valor: %d\n", i, v)
	}
	
}


func main() {
    arrays()
    iterating()
}
