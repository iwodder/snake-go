package main

import (
	"fmt"
	"slices"
	"snake/ui"
	"time"
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
	up direction = iota
	right
	down
	left
)

type snake struct {
	ui.SnakeComponent
	moveTimer      time.Duration
	moveDelay      time.Duration
	lastLength     int
	startingLength int
	dir            direction
}

func (s *snake) changeDirection(d direction) {
	if !d.isOpposite(s.dir) {
		s.dir = d
	}
}

func (s *snake) Update(g *game, delta time.Duration) {
	if !s.canMove(g.GameBoard, delta) {
		return
	}
	nextPos := s.head()
	switch s.dir {
	case up:
		nextPos.Y--
	case right:
		nextPos.X++
	case down:
		nextPos.Y++
	case left:
		nextPos.X--
	}
	s.Body = append(s.Body, nextPos)

	if s.crashed() {
		if g.remainingLives -= 1; g.remainingLives > 0 {
			s.ResetTo(g.GameBoard.Center())
		}
		return
	}

	if cnt := s.eat(g.apples); cnt > 0 {
		g.score += cnt * pointsPerApple
	}
	s.Body = s.Body[1:]
}

func (s *snake) canMove(b *ui.GameBoard, delta time.Duration) bool {
	s.moveTimer -= delta
	if s.moveTimer > 0 {
		return false
	}
	s.moveTimer = s.moveDelay

	c := s.head()
	return !((c.X >= b.Right() && s.dir == right) ||
		(c.X <= b.Left() && s.dir == left) ||
		(c.Y <= b.Top() && s.dir == up) ||
		(c.Y >= b.Bottom() && s.dir == down))
}

func (s *snake) eat(as apples) uint {
	ret := uint(0)
	p := s.headPos()
	for i := range as {
		if p == as[i].pos {
			as[i].eaten = true
			s.Body = slices.Insert(s.Body, 0, s.Body[0])
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
	s.lastLength = len(s.Body)
}

func (s *snake) shouldIncreaseSpeed() bool {
	return len(s.Body) >= s.lastLength*2
}

func (s *snake) headPos() ui.Position {
	head := s.head()
	return ui.Position{X: head.X, Y: head.Y}
}

func (s *snake) head() ui.Position {
	return s.Body[len(s.Body)-1]
}

func (s *snake) crashed() bool {
	head := s.headPos()
	for i := 0; i < len(s.Body)-2; i += 1 {
		if head.X == s.Body[i].X && head.Y == s.Body[i].Y {
			return true
		}
	}
	return false
}

func (s *snake) Notify(event Event) {
	switch event {
	case MoveDown:
		s.changeDirection(down)
	case MoveUp:
		s.changeDirection(up)
	case MoveRight:
		s.changeDirection(right)
	case MoveLeft:
		s.changeDirection(left)
	}
}

func (s *snake) ResetTo(initial ui.Position) {
	s.init(initial)
}

func (s *snake) Length() int {
	return len(s.Body)
}

func (s *snake) init(initial ui.Position) {
	body := make([]ui.Position, 0, 48)
	zeroBasedCol := initial.X - s.startingLength + 1
	for range s.startingLength {
		body = append(body, ui.Position{X: zeroBasedCol, Y: initial.Y})
		zeroBasedCol += 1
	}

	s.dir = startingDir
	s.moveDelay = defaultStartingSnakeMoveDelay
	s.lastLength = len(body)
	s.Body = body
}

func newSnake(initial ui.Position) *snake {
	return newSnakeOfLength(initial, DefaultStartingLength)
}

func newSnakeOfLength(initial ui.Position, length int) *snake {
	ret := snake{
		startingLength: length,
	}
	ret.init(initial)

	return &ret
}
