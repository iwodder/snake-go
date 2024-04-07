package main

import (
	"github.com/gdamore/tcell/v2"
)

type direction uint

func (d direction) rune() rune {
	switch d {
	case up, down:
		return '|'
	case right, left:
		return '-'
	}
	panic("unsupported direction")
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

func (b bounds) isInside(p pos) bool {
	return (b.upperLeft.x < p.x && p.x < b.lowerRight.x) && (b.upperLeft.y < p.y && p.y < b.lowerRight.y)
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
			case up:
				y--
			case left:
				x--
			case right:
				x++
			}
			scn.SetContent(x, y, m.dir.rune(), nil, s.style)
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
	m := &s.segments[0]
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
	if len(s.segments) < 2 {
		return
	}
	m.mag--
	s.segments[len(s.segments)-1].mag++
	if m.mag == 0 {
		s.segments = s.segments[1:]
	}
}

func (s *snake) eat(as apples) {
	p := s.headPos()
	for i := range as {
		if p == as[i].pos {
			as[i].eaten = true
			s.segments[len(s.segments)-1].mag++
		}
	}
}

func (s *snake) headPos() pos {
	ret := pos{x: s.start.x, y: s.start.y}
	for _, seg := range s.segments {
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

func newSnake(x int, y int) *snake {
	return &snake{
		style:    tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite),
		start:    pos{x, y},
		segments: append(make([]vector, 0, 24), vector{dir: right, mag: 1}),
	}
}
