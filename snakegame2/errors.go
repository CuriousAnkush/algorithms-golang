package snakegame2

import "fmt"

var ErrSnakeOnBoundary = fmt.Errorf("snake hit on the boundary")
var ErrSnakeCollidedWithBody = fmt.Errorf("snake collided with the body")
