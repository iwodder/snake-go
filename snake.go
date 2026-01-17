package main

import (
	"fmt"
	"slices"
	"snake/ui"
	"time"

	"github.com/gdamore/tcell/v2"
)

const (
	startingDir = right

	defaultStartingSnakeMoveDelay = time.Millisecond * 250
)

type direction uint

func (d direction) isOpposite(o direction) bool {
	return (d == right && o == left) ||
		(d == left && o == right) ||
		(d == up && o == down) ||
		(d == down && o == up)
}

func (d direction) String() string {
	switch d {
	case up:
		return "up"
	case down:
		return "down"
	case right:
		return "right"
	case left:
		return "left"
	default:
		return fmt.Sprintf("unrecognized direction: %d", d)
	}
}

const (
	snakeRune = 'X'

	up direction = iota
	right
	down
	left
)

type Position = ui.Position

type cell struct {
	x, y int
}

type snake struct {
	moveTimer      time.Duration
	moveDelay      time.Duration
	lastLength     int
	startingLength int
	dir            direction
	body           []cell
}

func (s *snake) draw(scn tcell.Screen) {
	for _, c := range slices.All(s.body) {
		scn.SetContent(c.x, c.y, snakeRune, nil, snakeStyle)
	}
}

func (s *snake) moveLeft() {
	s.changeDirection(left)
}

func (s *snake) moveRight() {
	s.changeDirection(right)
}

func (s *snake) moveUp() {
	s.changeDirection(up)
}

func (s *snake) moveDown() {
	s.changeDirection(down)
}

func (s *snake) changeDirection(d direction) {
	if !d.isOpposite(s.dir) {
		s.dir = d
	}
}

func (s *snake) move(b *board, delta time.Duration) {
	if !s.canMove(b, delta) {
		return
	}
	nextCell := cell{
		x: s.head().x,
		y: s.head().y,
	}
	switch s.dir {
	case up:
		nextCell.y--
	case right:
		nextCell.x++
	case down:
		nextCell.y++
	case left:
		nextCell.x--
	}
	s.body = append(s.body, nextCell)
	s.body = s.body[1:]
}

func (s *snake) canMove(b *board, delta time.Duration) bool {
	s.moveTimer -= delta
	if s.moveTimer > 0 {
		return false
	}
	s.moveTimer = s.moveDelay

	c := s.head()
	return !((c.x >= b.rightEdge() && s.dir == right) ||
		(c.x <= b.leftEdge() && s.dir == left) ||
		(c.y <= b.topEdge() && s.dir == up) ||
		(c.y >= b.bottomEdge() && s.dir == down))
}

func (s *snake) eat(as apples) uint {
	ret := uint(0)
	p := s.headPos()
	for i := range as {
		if p == as[i].pos {
			as[i].eaten = true
			s.body = slices.Insert(s.body, 0, s.body[0])
			ret += 1
		}
	}
	if s.shouldIncreaseSpeed() {
		s.speedUp()
	}
	return ret
}

func (s *snake) speedUp() {
	s.moveDelay = time.Duration(float64(s.moveDelay) * 0.75)
	s.lastLength = len(s.body)
}

func (s *snake) shouldIncreaseSpeed() bool {
	return len(s.body) >= s.lastLength*2
}

func (s *snake) headPos() Position {
	head := s.head()
	return Position{X: head.x, Y: head.y}
}

func (s *snake) head() cell {
	return s.body[len(s.body)-1]
}

func (s *snake) crashed() bool {
	head := s.headPos()
	for i := 0; i < len(s.body)-2; i += 1 {
		if head.X == s.body[i].x && head.Y == s.body[i].y {
			return true
		}
	}
	return false
}

func (s *snake) Notify(event Event) {
	switch event {
	case MoveDown:
		s.moveDown()
	case MoveUp:
		s.moveUp()
	case MoveRight:
		s.moveRight()
	case MoveLeft:
		s.moveLeft()
	}
}

func (s *snake) ResetTo(initial Position) {
	s.init(initial)
}

func (s *snake) Length() int {
	return len(s.body)
}

func (s *snake) init(initial Position) {
	body := make([]cell, 0, 48)
	zeroBasedCol := initial.X - s.startingLength + 1
	for range s.startingLength {
		body = append(body, cell{x: zeroBasedCol, y: initial.Y})
		zeroBasedCol += 1
	}

	s.dir = startingDir
	s.moveDelay = defaultStartingSnakeMoveDelay
	s.lastLength = len(body)
	s.body = body
}

func newSnake(initial Position) *snake {
	return newSnakeOfLength(initial, DefaultStartingLength)
}

func newSnakeOfLength(initial Position, length int) *snake {
	ret := snake{
		startingLength: length,
	}
	ret.init(initial)

	return &ret
}
