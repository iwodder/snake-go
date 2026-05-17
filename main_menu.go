package main

import (
	"snake/ui"
)

type MainMenu struct {
	*ui.GameBoardRenderer
}

func NewMainMenu(ul ui.Position, width int, height int) *MainMenu {
	boardRenderer := ui.NewGameBoardRenderer(ul, width, height)

	menuWidth := (width / 10) * 7
	padding := (boardRenderer.Width() - menuWidth) / 2

	x := boardRenderer.Left() + padding
	y := boardRenderer.Height() / 4

	mainMenu := ui.NewMenu(
		ui.Position{X: x, Y: y},
		menuWidth, height/2,
		"Main Menu",
	)
	mainMenu.AddEntry("")
	mainMenu.AddEntry("Enter to Start")
	mainMenu.AddEntry("SpcBr to Pause")
	mainMenu.AddEntry("Ctrl-C to Exit")

	_ = boardRenderer.Add(mainMenu)

	return &MainMenu{boardRenderer}
}
