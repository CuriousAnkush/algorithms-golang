package snakegame

type Snake struct {
	body  *List
	board *Board
}

func (s *Snake) moveHead(x int, y int) int {
	currentHead := s.body.head
	s.body.head = &Node{
		next:  nil,
		value: []int{x, y},
	}
	s.body.head.next = currentHead

	if s.board.FoodAvailable(s.body.head.value[0], s.body.head.value[1]) {
		s.board.EatFood()
		return 1
	}
	s.adjustLength(currentHead)
	if s.GotKilled() {
		return -1
	}
	return 0
}

func (s *Snake) adjustLength(currentHead *Node) {
	// When snake length is just one
	if currentHead == s.body.tail {
		s.body.head.next = nil
		s.body.tail = s.body.head
		return
	}
	s.trimTail()
}

func (s *Snake) trimTail() {
	head := s.body.head
	for head != nil && head.next != nil {
		if head.next.next == nil {
			head.next = nil
		}
		head = head.next
	}
}

func (s *Snake) Move(direction Direction) int {
	switch direction {
	case UP:
		return s.moveHead(s.body.head.value[0]-1, s.body.head.value[1])
	case Down:
		return s.moveHead(s.body.head.value[0]+1, s.body.head.value[1])
	case Right:
		return s.moveHead(s.body.head.value[0], s.body.head.value[1]+1)
	case Left:
		return s.moveHead(s.body.head.value[0], s.body.head.value[1]-1)
	}
	return 0
}

func (s *Snake) OnBoundary() bool {
	if s.body.head.value[0] < 0 || s.body.head.value[0] >= s.board.height {
		return true
	}
	if s.body.head.value[1] < 0 || s.body.head.value[1] >= s.board.widht {
		return true
	}
	if s.body.tail.value[0] < 0 || s.body.tail.value[0] >= s.board.height {
		return true
	}
	if s.body.tail.value[1] < 0 || s.body.tail.value[1] >= s.board.widht {
		return true
	}

	return false
}

func (s *Snake) HeadCollidedWithBody() bool {
	currentHeadNext := s.body.head.next

	for currentHeadNext != nil && currentHeadNext.next != nil {
		if currentHeadNext.value[0] == s.body.head.value[0] && currentHeadNext.value[1] == s.body.head.value[1] {
			return true
		}
		currentHeadNext = currentHeadNext.next
	}
	return false
}

func (s *Snake) GotKilled() bool {
	return s.HeadCollidedWithBody() || s.OnBoundary()
}

func SnakeConstructor(b *Board) *Snake {
	initialPos := &Node{
		next:  nil,
		value: []int{0, 0},
	}
	return &Snake{body: &List{
		head: initialPos,
		tail: initialPos,
	},
		board: b,
	}
}

type List struct {
	head *Node
	tail *Node
}

type Node struct {
	next  *Node
	value []int
}
