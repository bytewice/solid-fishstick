package reidemeister

// Estrutura para representar um nó
type Knot struct {
	Crossings int      // Número de cruzamentos
	Matrix    [][]int  // Matriz de Seifert
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

// Função para deslizar duas curvas (Movimento II)
func reidemeisterII(knot Knot, i, j int) Knot {
	n := knot.Crossings

	// Verifica se os índices são válidos
	if i < 0 || i >= n || j < 0 || j >= n || i == j {
		panic("Índices inválidos para o Movimento II.")
	}

	// Cria uma nova matriz com uma linha e coluna a menos
	newMatrix := make([][]int, n-1)
	for k := range newMatrix {
		newMatrix[k] = make([]int, n-1)
	}

	// Combina as curvas i e j
	for k := 0; k < n-1; k++ {
		for l := 0; l < n-1; l++ {
			if k < i && l < i {
				newMatrix[k][l] = knot.Matrix[k][l]
			} else if k >= i && l < i {
				newMatrix[k][l] = knot.Matrix[k+1][l]
			} else if k < i && l >= i {
				newMatrix[k][l] = knot.Matrix[k][l+1]
			} else {
				newMatrix[k][l] = knot.Matrix[k+1][l+1]
			}
		}
	}

	return Knot{
		Crossings: n - 1,
		Matrix:    newMatrix,
	}
}

// Função para deslizar uma curva sobre um cruzamento (Movimento III)
func reidemeisterIII(knot Knot, i, j int) Knot {
	n := knot.Crossings

	// Verifica se os índices são válidos
	if i < 0 || i >= n || j < 0 || j >= n || i == j {
		panic("Índices inválidos para o Movimento III.")
	}

	// Cria uma cópia da matriz original
	newMatrix := make([][]int, n)
	for k := range newMatrix {
		newMatrix[k] = make([]int, n)
		copy(newMatrix[k], knot.Matrix[k])
	}

	// Permuta as linhas i e j
	newMatrix[i], newMatrix[j] = newMatrix[j], newMatrix[i]

	// Permuta as colunas i e j
	for k := 0; k < n; k++ {
		newMatrix[k][i], newMatrix[k][j] = newMatrix[k][j], newMatrix[k][i]
	}

	return Knot{
		Crossings: n,
		Matrix:    newMatrix,
	}
}
