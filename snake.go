package main

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
)

var dirRunes = map[direction]rune{
	up: tcell.RuneVLine, down: tcell.RuneVLine, right: tcell.RuneHLine, left: tcell.RuneHLine,
}

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
	up = iota
	right
	down
	left
)

type pos struct {
	x, y int
}

type boundary struct {
	upperLeft  pos
	lowerRight pos
}

// isInside reports whether the argument is within the boundary. The boundary are adjusted to be zero-based.
func (b boundary) isInside(p pos) bool {
	return (b.upperLeft.x < p.x && p.x < b.lowerRight.x-1) && (b.upperLeft.y < p.y && p.y < b.lowerRight.y-1)
}

func (b boundary) leftEdge() int {
	return b.upperLeft.x
}

func (b boundary) rightEdge() int {
	return b.lowerRight.x - 1
}

func (b boundary) topEdge() int {
	return b.upperLeft.y
}

func (b boundary) bottomEdge() int {
	return b.lowerRight.y - 1
}

func (b boundary) height() int {
	return b.lowerRight.y - b.upperLeft.y
}

func (b boundary) width() int {
	return b.lowerRight.x - b.upperLeft.x
}

func (b boundary) shrink(amt int) boundary {
	return boundary{
		upperLeft: pos{
			x: b.upperLeft.x + amt,
			y: b.upperLeft.y + amt,
		},
		lowerRight: pos{
			x: b.lowerRight.x - amt,
			y: b.lowerRight.y - amt,
		},
	}
}

func (b boundary) center() pos {
	return pos{x: b.width() / 2, y: b.height() / 2}
}

type vector struct {
	dir direction
	mag int
	r   rune
}

func (v vector) draw(scn tcell.Screen, start pos, style tcell.Style) pos {
	for range v.mag {
		switch v.dir {
		case up:
			start.y--
		case down:
			start.y++
		case right:
			start.x++
		case left:
			start.x--
		}
		scn.SetContent(start.x, start.y, v.r, nil, style)
	}
	return start
}

type snake struct {
	start pos
	vecs  []vector
	timer time.Duration
}

func (s *snake) draw(scn tcell.Screen) {
	currentPos := s.start
	for idx, vec := range s.vecs {
		currentPos = vec.draw(scn, currentPos, snakeStyle)
		if idx < len(s.vecs)-1 {
			scn.SetContent(currentPos.x, currentPos.y, s.getRune(vec.dir, s.vecs[idx+1].dir), nil, snakeStyle)
		}
	}
}

func (s *snake) getRune(curDir direction, nextDir direction) rune {
	switch {
	case (curDir == up || curDir == left) && (nextDir == right || nextDir == down):
		return tcell.RuneULCorner
	case (curDir == up || curDir == right) && (nextDir == left || nextDir == down):
		return tcell.RuneURCorner
	case (curDir == down || curDir == left) && (nextDir == right || nextDir == up):
		return tcell.RuneLLCorner
	case (curDir == down || curDir == right) && (nextDir == left || nextDir == up):
		return tcell.RuneLRCorner
	default:
		panic(fmt.Sprintf("unhandled direction combination: curDir=%s, nextDir=%s", curDir, nextDir))
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
	if last := s.head(); s.isNewDirectionValid(last.dir, d) {
		if last.mag == 0 {
			last.dir = d
		} else {
			s.vecs = append(s.vecs, vector{dir: d, mag: 0, r: dirRunes[d]})
		}
	}
}

func (s *snake) isNewDirectionValid(last, new direction) bool {
	return last != new && !new.isOpposite(last)
}

func (s *snake) move(b boundary, delta time.Duration) {
	if !s.canMove(b, delta) {
		return
	}
	switch s.tail().dir {
	case up:
		s.start.y--
	case right:
		s.start.x++
	case down:
		s.start.y++
	case left:
		s.start.x--
	}
	if len(s.vecs) < 2 {
		return
	}
	s.tail().mag--
	s.head().mag++
	if s.tail().mag == 0 {
		s.vecs = s.vecs[1:]
	}
}

func (s *snake) canMove(b boundary, delta time.Duration) bool {
	s.timer -= delta
	if s.timer > 0 {
		return false
	} else {
		s.timer = time.Millisecond * 250
	}

	p := s.headPos()
	if b.isInside(p) {
		return true
	}
	lastDir := s.head().dir
	return !((p.x >= b.rightEdge() && lastDir == right) ||
		(p.x <= b.leftEdge() && lastDir == left) ||
		(p.y <= b.topEdge() && lastDir == up) ||
		(p.y >= b.bottomEdge() && lastDir == down))
}

func (s *snake) eat(as apples) {
	p := s.headPos()
	for i := range as {
		if p == as[i].pos {
			as[i].eaten = true
			s.vecs[len(s.vecs)-1].mag++
		}
	}
}

func (s *snake) headPos() pos {
	ret := pos{x: s.start.x, y: s.start.y}
	for _, seg := range s.vecs {
		switch seg.dir {
		case up:
			ret.y -= seg.mag
		case down:
			ret.y += seg.mag
		case left:
			ret.x -= seg.mag
		case right:
			ret.x += seg.mag
		}
	}
	return ret
}

func (s *snake) head() *vector {
	return &s.vecs[len(s.vecs)-1]
}

func (s *snake) tail() *vector {
	return &s.vecs[0]
}

func (s *snake) notify(ev *tcell.EventKey) {
	switch ev.Key() {
	case tcell.KeyDown:
		s.moveDown()
	case tcell.KeyUp:
		s.moveUp()
	case tcell.KeyRight:
		s.moveRight()
	case tcell.KeyLeft:
		s.moveLeft()
	}
}

func newSnake(initial pos) *snake {
	return &snake{
		start: initial,
		vecs:  append(make([]vector, 0, 24), vector{dir: right, mag: 1, r: dirRunes[right]}),
	}
}
