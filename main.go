package main

import (
	"github.com/gdamore/tcell/v2"
	"log"
	"os"
)

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
