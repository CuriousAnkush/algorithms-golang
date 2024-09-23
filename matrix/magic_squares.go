package matrix

func numMagicSquaresInside(grid [][]int) int {
	var countMatrix int
	gridHeight := len(grid)
	gridWidth := len(grid[0])

	if gridWidth < 3 || gridHeight < 3 {
		return 0
	}

	for row := 0; row < gridHeight-2; row++ {
		for col := 0; col < gridWidth-2; col++ {
			if !validMatrix(&grid, row, col) {
				continue
			}
			countMatrix++
		}
	}

	return countMatrix
}

func validMatrix(i *[][]int, row, col int) bool {
	matrix := *i
	set := make(map[int]bool)
	var rowSum [3]int
	var colSum [3]int

	firstDiagonalSum := matrix[row][col] + matrix[row+1][col+1] + matrix[row+2][col+2]
	secondDiagonalSum := matrix[row][col+2] + matrix[row+1][col+1] + matrix[row+2][col]

	for r := row; r < row+3; r++ {
		for c := col; c < col+3; c++ {
			rowSum[r-row] += matrix[r][c]
			colSum[c-col] += matrix[r][c]
			if matrix[r][c] > 9 {
				return false
			}

			if set[matrix[r][c]] {
				return false
			}

			set[matrix[r][c]] = true
		}
	}

	sum := rowSum[0]

	if sum == rowSum[1] && sum == rowSum[2] && sum == colSum[0] &&
		sum == colSum[1] && sum == colSum[2] && sum == firstDiagonalSum && sum == secondDiagonalSum {
		return true
	}

	return false
}
