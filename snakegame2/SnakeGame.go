package snakegame2

type SnakeGame struct {
	board        *Board
	snake        *Snake
	currentScore int
}

func Constructor(width int, height int, food [][]int) SnakeGame {
	board := NewBoard(width, height, food)
	snake := NewSnake(board)
	return SnakeGame{
		board: board,
		snake: snake,
	}

}

func (g *SnakeGame) Move(direction string) int {
	score := g.snake.Move(Direction(direction))
	if score == -1 {
		return score
	}
	g.currentScore += score
	return g.currentScore
}
