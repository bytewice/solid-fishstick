package main

import "fmt"

func main(){
	var x int32 = 1
	var y int16 = 2

	// Não pode fazer isso porque as duas variáveis são de tipos diferentes
	// x = y

	// jeito correto de fazer essa conversão, bem parecido com C também
	x = int32(y)
}

func float(){
	// float32 ~6 digítos de precisão
	// float64 ~15 digítos de precisão 
	var x float64 = 123.45 
}

func complex(){
	// z = 2 + 3i
	var z complex128 = complex(2,3)
}

// Unicode é pra representar todos os caracteres que existem aparentemente
// UTF-8 tem entre 8 e 32 bits de comprimento