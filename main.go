package main

import (
	"github.com/gdamore/tcell/v2"
	"log"
	"os"
)

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
