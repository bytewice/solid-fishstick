package main

import (
	"os"
	"fmt"
)

func main(){

	// abre o arquivo
	// os.Open("outfile.txt")

	//fecha o arquivo
	// os.CLose("outfile.txt")

	// lê o arquivo num array de bytes
	// os.Read("outfile.txt")

	// costumava ser f, err := os.Open .......
	f, _ := os.Open("outfile.txt")
	// lê só os 10 primeiros bytes
	barr := make([]byte, 10)
	nb, _ := f.Read(barr)
	f.Close()

	_ = nb
	fmt.Println(string(barr))

	f, err := os.Create("outfile2.txt")

	barr2 := []byte{1, 2, 3}
	// pode ser um ou outro
	//nc, err := f.Write(barr2)
	nc, err := f.WriteString("Hi")

	 _ = err
	 _ = nc
	 _ = barr2
}