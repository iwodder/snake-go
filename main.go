package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

func main() {
	scn, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("failed to get screen: %v", err)
	}
	if err = scn.Init(); err != nil {
		log.Fatalf("failed to init screen: %v", err)
	}
	err = RunGame(newGame(scn), scn)
	scn.Fini()
	if err != nil {
		log.Fatalf("error while running game: %v", err)
	}
}
