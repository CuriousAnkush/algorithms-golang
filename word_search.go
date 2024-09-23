package algorithms

//Given an m x n grid of characters board and a string word, return true if word exists in the grid.
//
//The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or
//vertically neighboring. The same letter cell may not be used more than once.
//board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
//Output: true

func exist(board [][]string, word string) bool {
	for i := range word {
		for row := range board {
			for column := range board[row] {
				if string(word[i]) != board[row][column] {
					continue
				}
				visited := make([][]bool, len(board))
				for i := 0; i < len(board); i++ {
					visited[i] = make([]bool, len(board[0]))
				}

				if dfs(word[i+1:], board, row, column, visited) {
					return true
				}
			}
		}
		return false
	}

	return false
}

func dfs(word string, board [][]string, row int, col int, visited [][]bool) bool {
	if len(word) == 0 {
		return true
	}
	if visited[row][col] {
		return false
	}
	visited[row][col] = true

	directions := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

	for i := range directions {
		ro := row + directions[i][0]
		co := col + directions[i][1]
		if ro < 0 || ro >= len(board) || co < 0 || co >= len(board[0]) || visited[ro][co] || string(word[0]) != board[ro][co] {
			continue
		}
		if dfs(word[1:], board, ro, co, visited) {
			return true
		}

	}
	visited[row][col] = false

	return false
}
