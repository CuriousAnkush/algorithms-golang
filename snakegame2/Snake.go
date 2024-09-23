package snakegame2

import (
	"container/list"
)

type Snake struct {
	body      *list.List
	board     *Board
	bodyCells map[BodyCell]bool
}

type BodyCell struct {
	row int
	col int
}

func NewSnake(board *Board) *Snake {
	body := list.New()
	body.PushFront([]int{0, 0})
	return &Snake{body: body,
		board:     board,
		bodyCells: make(map[BodyCell]bool)}
}

func (s Snake) Move(direction Direction) int {
	var status error
	var foodFound bool
	snakeHeadLocation := s.body.Front().Value.([]int)
	switch direction {
	case UP:
		foodFound, status = s.MoveHead(snakeHeadLocation[0]-1, snakeHeadLocation[1])
	case DOWN:

		foodFound, status = s.MoveHead(snakeHeadLocation[0]+1, snakeHeadLocation[1])
	case Left:
		foodFound, status = s.MoveHead(snakeHeadLocation[0], snakeHeadLocation[1]-1)
	case Right:
		foodFound, status = s.MoveHead(snakeHeadLocation[0], snakeHeadLocation[1]+1)
	}
	if status != nil {
		return -1
	}

	if foodFound {
		return 1
	} else {
		return 0
	}

}

func (s Snake) IsHeadCollidedWithBody() bool {
	if s.body.Len() == 1 || s.body.Len() == 2 {
		return false
	}
	head := s.body.Front()

	nextBodyCell := s.body.Front().Next()

	for nextBodyCell.Next() != nil && nextBodyCell.Next().Value != nil {
		headValM := head.Value.([]int)
		if _, found := s.bodyCells[BodyCell{
			row: headValM[0],
			col: headValM[1],
		}]; found {
			return true
		}
		nextBodyCell = nextBodyCell.Next()
	}

	return false
}

func (s Snake) MoveHead(rowNu int, colNu int) (bool, error) {
	if s.board.IsBoundary(rowNu, colNu) {
		return false, ErrSnakeOnBoundary
	}

	lastHead := s.body.Front()
	s.body.PushFront([]int{rowNu, colNu})

	if lastHead != s.body.Back() {
		lastHeadVal := lastHead.Value.([]int)
		s.bodyCells[BodyCell{
			row: lastHeadVal[0],
			col: lastHeadVal[1],
		}] = true
	}

	if s.board.FoodAvailable(rowNu, colNu) {
		s.board.EatFood()
		return true, nil
	}

	s.adjustLength()
	if s.IsHeadCollidedWithBody() {
		return false, ErrSnakeCollidedWithBody
	}
	return false, nil
}

func (s Snake) adjustLength() {
	tail := s.body.Back()
	s.body.Remove(tail)
	cellVal := tail.Value
	cellValM := cellVal.([]int)
	delete(s.bodyCells, BodyCell{
		row: cellValM[0],
		col: cellValM[1],
	})
}
