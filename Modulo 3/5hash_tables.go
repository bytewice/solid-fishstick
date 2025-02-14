package main

import "fmt"

// Maps are the implementaion of a hash table
func maps(){
	//        map[ 'key type' ] 'value type'
	var idMap map[string]int
	idMap = make(map[string]int)

	idMap := map[string]int {
		"joe": 123
	}

	// acessando maps
	fmt.Println(idMap["joe"])
	// retorna 0 se a chave não existir

	// adicionando ou atualizando um par de chave/valor no mapa
	idMap["jane"] = 456

	// deletando um par do mapa
	delete(idMap, "joe")

	// MAP FUNCTIONS

	// p retorna false se joe o 'id' não estiver no mapa
	id, p := idMap["joe"]

	// Retorna quantos pares estão no mapa
	fmt.Println(len(idMap))

	for key, val := range idMap {
		fmt.Println(key, val)
	}
}