package main

import (
	"github.com/gdamore/tcell/v2"
	"log"
)

func main() {
	scn, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("failed to get screen: %v", err)
	}
	if err = scn.Init(); err != nil {
		log.Fatalf("failed to init screen: %v", err)
	}
	newGame(scn).start()
}
