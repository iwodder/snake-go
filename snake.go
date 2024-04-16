package main

import (
	"github.com/gdamore/tcell/v2"
)

var dirRunes = map[direction]rune{
	up: '|', down: '|', right: '-', left: '-',
}

type direction uint

func (d direction) isOpposite(o direction) bool {
	return (d == right && o == left) ||
		(d == left && o == right) ||
		(d == up && o == down) ||
		(d == down && o == up)
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

type bounds struct {
	upperLeft  pos
	lowerRight pos
}

// isInside reports whether the argument is within the bounds. The bounds are adjusted to be zero-based.
func (b bounds) isInside(p pos) bool {
	return (b.upperLeft.x < p.x && p.x < b.lowerRight.x-1) && (b.upperLeft.y < p.y && p.y < b.lowerRight.y-1)
}

func (b bounds) leftEdge() int {
	return b.upperLeft.x
}

func (b bounds) rightEdge() int {
	return b.lowerRight.x - 1
}

func (b bounds) topEdge() int {
	return b.upperLeft.y
}

func (b bounds) bottomEdge() int {
	return b.lowerRight.y - 1
}

func (b bounds) height() int {
	return b.lowerRight.y - b.upperLeft.y
}

func (b bounds) width() int {
	return b.lowerRight.x - b.upperLeft.x
}

func (b bounds) shrink(amt int) bounds {
	return bounds{
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
	style tcell.Style
	start pos
	vecs  []vector
}

func (s *snake) draw(scn tcell.Screen) {
	p := s.start
	for _, m := range s.vecs {
		p = m.draw(scn, p, s.style)
	}
}

func (s *snake) headLeft() {
	s.changeDirection(left)
}

func (s *snake) headRight() {
	s.changeDirection(right)
}

func (s *snake) headUp() {
	s.changeDirection(up)
}

func (s *snake) headDown() {
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

func (s *snake) move(b bounds) {
	if !s.canMove(b) {
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

func (s *snake) canMove(b bounds) bool {
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

func newSnake(x int, y int) *snake {
	return &snake{
		style: tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite),
		start: pos{x, y},
		vecs:  append(make([]vector, 0, 24), vector{dir: right, mag: 1, r: dirRunes[right]}),
	}
}
