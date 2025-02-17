package main

import {
	"fmt"
	"ioutil"
	"io"
	"os"
}

// ioutil
func main(){
	dat, e :=  ioutil.ReadFile("text.txt")

	// tem que converter pra string
	fmt.Println(string(dat))
	//caso não converte ele printa os valores em bytes
	fmt.Println(dat)

	_ = e

	// creating a new file
	data := []byte("Hello, world!")
	err := ioutil.WriteFile("outfile.txt", data, 0777)
	// aparentemente ioutil foi descontinuado e as funções dessa biblioteca foram movidas para os pacotes os e io
	_ = err
}