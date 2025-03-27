package main

import (
	"fmt"
//	"my_project/reidemeister" // Importa o pacote criado
)

type Knot struct{
	Crossings int //cruzamentos
	Matrix [][]int //matriz de seifert
}

// generateSeifertMatrix generates the Seifert matrix of a knot with n crossings.
func generateSeifertMatrix(n int) [][]int {
	if n <= 0 {
		panic("The number of crossings must be greater than zero.")
	}

	// Initialize an n x n matrix filled with zeros
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	// Fill the Seifert matrix
	for i := 0; i < n; i++ {
		matrix[i][i] = 2 // Main diagonal
		if i < n-1 {
			matrix[i][i+1] = -1 // Adjacent elements
			matrix[i+1][i] = -1 // Symmetry of the matrix
		}
	}

	return matrix
}

func createKnot(crossings int) Knot{
	return Knot{
			Crossings: crossings,
			Matrix: generateSeifertMatrix(crossings),
	}
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}

func generateTrefoil() [][]int {
	return [][]int{
		{ 1, -1, 0},
		{-1,  1,-1},
		{ 0, -1, 1},
	}
}

// Função para adicionar uma volta (Movimento I)
func reidemeisterI(knot Knot) Knot {
	n := knot.Crossings

	// Cria uma nova matriz com uma linha e coluna adicionais
	newMatrix := make([][]int, n+1)
	for i := range newMatrix {
		newMatrix[i] = make([]int, n+1)
	}

	// Copia a matriz original para a nova matriz
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			newMatrix[i][j] = knot.Matrix[i][j]
		}
	}

	// Adiciona a nova curva de Seifert
	newMatrix[n][n] = 2 // Diagonal principal
	if n > 0 {
		newMatrix[n][n-1] = -1 // Conexão com a curva anterior
		newMatrix[n-1][n] = -1 // Simetria da matriz
	}

	return Knot{
		Crossings: n + 1,
		Matrix:    newMatrix,
	}
}

func main() {
	nCrossings := 4 // Number of crossings in the knot

	seifertMatrix := generateSeifertMatrix(nCrossings)

	trefoil := Knot{
		Crossings: 3,
		Matrix: generateTrefoil(),
	}

	fmt.Println("Seifert Matrix:")
	printMatrix(seifertMatrix)
	
	fmt.Println("Trefoil:")
	printMatrix(trefoil.Matrix)

	fmt.Println("Reidemeister I no trefoil...")
	trefoil = reidemeisterI(trefoil)
	printMatrix(trefoil.Matrix)
}
