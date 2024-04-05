package main

import "github.com/gdamore/tcell/v2"

type direction uint

const (
	up = iota
	right
	down
	left
)

type pos struct {
	x, y int
}

type move struct {
	d         direction
	magnitude int
}

type snake struct {
	style    tcell.Style
	start    pos
	movement []move
}

func (s *snake) draw(scn tcell.Screen) {
	x, y := s.start.x, s.start.y
	for _, m := range s.movement {
		for range m.magnitude {
			switch m.d {
			case down:
				y++
				scn.SetContent(y, x, '|', nil, s.style)
			case up:
				y--
				scn.SetContent(y, x, '|', nil, s.style)
			case left:
				x--
				scn.SetContent(y, x, '-', nil, s.style)
			case right:
				x++
				scn.SetContent(y, x, '-', nil, s.style)
			}
		}
	}
}

func (s *snake) left() {
	s.move(left)
}

func (s *snake) right() {
	s.move(right)
}

func (s *snake) up() {
	s.move(up)
}

func (s *snake) down() {
	s.move(down)
}

func (s *snake) move(d direction) {
	if l := len(s.movement); l == 0 || s.movement[l-1].d != d {
		s.movement = append(s.movement, move{d: d, magnitude: 0})
	}
}

func (s *snake) updatePosition() {
	m := s.movement[0]
	switch m.d {
	case up:
		s.start.y--
	case right:
		s.start.x++
	case down:
		s.start.y++
	case left:
		s.start.x--
	}
	s.movement[len(s.movement)-1].magnitude++
}
