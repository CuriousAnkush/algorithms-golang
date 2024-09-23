package snakegame2

type Board struct {
	width      int
	height     int
	food       [][]int
	foodCursor int
}

type Direction string

const (
	UP    Direction = "U"
	DOWN  Direction = "D"
	Right Direction = "R"
	Left  Direction = "L"
)

func NewBoard(width int, height int, food [][]int) *Board {
	return &Board{
		width:      width,
		height:     height,
		food:       food,
		foodCursor: 0,
	}
}

func (b *Board) FoodAvailable(rowNu int, colNu int) bool {
	availableFoodCell := b.food[b.foodCursor]

	if b.foodCursor >= len(b.food) || rowNu != availableFoodCell[0] || colNu != availableFoodCell[1] {
		return false
	}

	return true
}

func (b *Board) EatFood() {
	b.foodCursor++
}

func (b *Board) IsBoundary(rowNu int, colNu int) bool {
	if rowNu >= b.height || colNu >= b.width || rowNu < 0 || colNu < 0 {
		return true
	}

	return false
}
