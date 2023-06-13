func equalPairs(grid [][]int) int {
	out := 0
	rev := reverseMatrix(grid)

	for i := range grid {
		for j := range grid {
			if reflect.DeepEqual(rev[i], grid[j]) {
				out++
			}
		}
	}
	return out
}

func reverseMatrix(matrix [][]int) [][]int {
	// Get the dimensions of the matrix
	rows := len(matrix)
	if rows == 0 {
		return matrix
	}
	cols := len(matrix[0])

	// Create a new matrix with reversed dimensions
	reversedMatrix := make([][]int, cols)
	for i := range reversedMatrix {
		reversedMatrix[i] = make([]int, rows)
	}

	// Traverse the original matrix and copy values to the reversed matrix
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			reversedMatrix[j][i] = matrix[i][j]
		}
	}

	return reversedMatrix
}
