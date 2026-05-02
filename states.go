package main

import "time"

const (
	GameOverText            = "Game Over"
	GamePausedText          = "Game Paused"
	MainMenuTransitionDelay = 2 * time.Second
)

type state interface {
	update(*game, time.Duration)
	handle(*game, Event)
}

type playingState struct {
	board *gameBoard
}

func (p *playingState) update(g *game, delta time.Duration) {
	p.board.Update(g, delta)
	if g.gameOver() {
		g.currentState = &gameOverState{delay: MainMenuTransitionDelay}
	}
}

func (p *playingState) handle(g *game, event Event) {
	if event == PauseGame {
		g.currentState = &pausedState{currentGame: p}
	}
}

type gameOverState struct {
	delay time.Duration
}

func (gos *gameOverState) update(g *game, delta time.Duration) {
	g.Manager.ShowModal(GameOverText)
	if gos.delay -= delta; gos.delay <= 0 {
		g.Manager.HideModal()
		g.currentState = new(menuState)
	}
}

func (gos *gameOverState) handle(*game, Event) {
	// do nothing
}

type pausedState struct {
	currentGame *playingState
}

func (p *pausedState) update(g *game, _ time.Duration) {
	g.Manager.ShowModal(GamePausedText)
}

func (p *pausedState) handle(g *game, event Event) {
	if event == PauseGame {
		g.Manager.HideModal()
		g.currentState = p.currentGame
	}
}

type menuState struct{}

func (m *menuState) update(g *game, _ time.Duration) {
	g.Manager.ShowModal("Press enter to start...")
}

func (m *menuState) handle(g *game, event Event) {
	if event == StartGame {
		g.reset()
		g.Manager.HideModal()
		g.currentState = &playingState{board: g.gameBoard}
	}
}
