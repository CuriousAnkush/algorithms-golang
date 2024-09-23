package snakegame

type Direction string

const (
	UP    Direction = "U"
	Down  Direction = "D"
	Right Direction = "R"
	Left  Direction = "L"
)

type SnakeGame struct {
	board        *Board
	snake        *Snake
	currentScore int
}

func Constructor(width int, height int, food [][]int) SnakeGame {
	board := NewBoard(height, width, food)
	snake := SnakeConstructor(board)
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

/**
 * Your SnakeGame object will be instantiated and called as such:
 * obj := Constructor(width, height, food);
 * param_1 := obj.Move(direction);
 */
