package strings

import "fmt"

// strings package
func string(){
	
	var s string
	// retorna uma cópia de s com as primeiras n ocorrencias de string(old) trocadas por string(new)
	x := Replace(s, old, new, n) 

	// bota todas as letras em maiúsculo
	ToLower(s)

	// bota todas as letras em minúsculo
	ToUpper(s)
}

func managing_strings (){
	// não é possível acessar diretamente o conteúdo de strings - elas são imutáveis
	// para fazer isso precisamos convertê-las para bytes e alterar diretamente os caracteres

	s:="Hello!"
	//convertendo de slices para bytes
	b:=s.[]bytes(s)
	// agr podemos modificar o que quisermos
	b[0] = 'M'
	// convertendo de volta para string
	newS:= string(b)
	fmt.Printf("%s", newS)
}

// strconv package
func strconv(){
	var s string = "12"

	// retorna o valor int se estiver escrito um número na string
	var num int = Atoi(s)

	// converte uma string para um número float
	ParseFloat(s, bitSize)
}