package snakegame

type Board struct {
	height     int
	widht      int
	food       [][]int
	foodCursor int
}

func NewBoard(height, width int, food [][]int) *Board {
	return &Board{
		height:     height,
		widht:      width,
		food:       food,
		foodCursor: 0,
	}
}
func (b *Board) FoodAvailable(x int, j int) bool {
	if b.foodCursor >= len(b.food) {
		return false
	}

	foodCell := b.food[b.foodCursor]
	if foodCell[0] == x && foodCell[1] == j {
		return true
	}

	return false
}

func (b *Board) EatFood() {
	b.foodCursor++
}
