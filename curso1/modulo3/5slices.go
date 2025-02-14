
import "fmt"

func main(){

	// Propriedades de um slice
	// Ponteiro : ponteiro apontando para o início do slice
	// Tamanho : tamanho do slice
	// Capacidade : é o tamanho máximo que o slice pode ter (( os slices podem variar de tamanho ))

	arr := [...]string{"a", "b", "c", "d", "e", "f", "g"}
	// s1 inclui arr[1] e arr[2], NÃO INCLUI arr[3]
	s1 := arr[1:3]
	// s2 inclui arr[2], arr[3], arr[4], NÃO INCLUI arr[5]
	s2 := arr[2:5]

	fmt.Printf(s1[1], s2[0])
	//O resultado vai ser o mesmo elemento, já que s1[1] = s2[0] = arr[2]

}

func variable_slices(){
	// 2 arguments make : type, capacity
	sli = make([]int, 10)
	// 3 arguments make : type, lenght, capacity
	sli_ = make([]int, 10, 15)

	// bota um elemento no final do slice, ele pode ultrapassar a capacidade do array, vai criando outro underlying array p suprir as necessidades
	_sli = make([]int, 0, 3)
	// tamanho do slice é 0
	_sli = append(sli, 100)
	// append aumenta a lenght do slice para um e insere '100'
}