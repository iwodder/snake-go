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

type vector struct {
	dir direction
	mag int
}

type snake struct {
	style    tcell.Style
	start    pos
	segments []vector
}

func (s *snake) draw(scn tcell.Screen) {
	x, y := s.start.x, s.start.y
	for _, m := range s.segments {
		for range m.mag {
			switch m.dir {
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
	if l := len(s.segments); l == 0 || s.segments[l-1].dir != d {
		s.segments = append(s.segments, vector{dir: d, mag: 0})
	}
}

func (s *snake) updatePosition() {
	m := s.segments[0]
	switch m.dir {
	case up:
		s.start.y--
	case right:
		s.start.x++
	case down:
		s.start.y++
	case left:
		s.start.x--
	}
	s.segments[len(s.segments)-1].mag++
}
