package main

import (
	"fmt"
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
}
