package main

import(
	"fmt"
)

// aparentemente p passar um slice faz name_slice []int
func getMax( vals ...int) int{
	maxV := -1
	for _, v := range vals{
		if v > maxV { maxV = v}
	}
	return maxV
}

func main(){

	// só vai executar depois que a função chegar ao fim --- nesse caso depois do fmt.Println("Hello!") no fim da main
	defer fmt.Println("\nBye!")

	sli := make([]int, 10)
	
	// 		esse loop cria uma cópia por debaixo dos panos
	// 		ele faz v = sli[i]
	// for i, v := range sli{
	// 	v = i
	// 	fmt.Println(v)
	// }

	// para alterarmos o valor do slice, teremos que fazer assim:
	for i := range sli{
		sli[i] = i
		fmt.Println(i)
	}
	
	fmt.Println(sli)

	// printar o getMax"
	fmt.Printf("valor máximo é %d\n",getMax(sli...))

	fmt.Println("Hello!")
}