package main

import (
	"github.com/gdamore/tcell/v2"
	"log"
	"os"
)

const NumCols = 80
const NumRows = 80

type frame [NumCols][NumRows]rune

func (f frame) render(s tcell.Screen) {
	for x := range NumCols {
		for y := range NumRows {
			s.SetContent(x, y, f[x][y], nil, tcell.Style{})
		}
	}
}

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
	start    pos
	movement []move
}

func (s *snake) draw(f *frame) {
	x, y := s.start.x, s.start.y
	for _, m := range s.movement {
		for range m.magnitude {
			switch m.d {
			case down:
				y++
				f[y][x] = '|'
			case up:
				y--
				f[y][x] = '|'
			case left:
				x--
				f[y][x] = '-'
			case right:
				x++
				f[y][x] = '-'
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

func main() {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalln(err)
	}
	err = s.Init()
	if err != nil {
		log.Fatalln(err)
	}

	for {
		// Poll event
		ev := s.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				s.Fini()
				os.Exit(0)
			}
		}
	}
}
