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
	var r rune
	for _, m := range s.segments {
		for range m.mag {
			switch m.dir {
			case down:
				r = '|'
				y++
			case up:
				r = '|'
				y--
			case left:
				r = '-'
				x--
			case right:
				r = '-'
				x++
			}
			scn.SetContent(x, y, r, nil, s.style)
		}
	}
}

func (s *snake) headLeft() {
	s.head(left)
}

func (s *snake) headRight() {
	s.head(right)
}

func (s *snake) headUp() {
	s.head(up)
}

func (s *snake) headDown() {
	s.head(down)
}

func (s *snake) head(d direction) {
	if l := len(s.segments); l == 0 || s.segments[l-1].dir != d {
		s.segments = append(s.segments, vector{dir: d, mag: 0})
	}
}

func (s *snake) move() {
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

func newSnake(x int, y int) *snake {
	return &snake{
		style:    tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite),
		start:    pos{x, y},
		segments: append(make([]vector, 0, 24), vector{dir: right, mag: 1}),
	}
}
