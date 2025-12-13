package main

import "github.com/gdamore/tcell/v2"

type Event int

const (
	Unknown Event = iota
	MoveUp
	MoveDown
	MoveLeft
	MoveRight
	PauseGame
	ExitGame
	StartGame
)

type EventListener interface {
	Notify(event Event)
}

type EventListeners []EventListener

func (e EventListeners) Notify(event Event) {
	for _, listener := range e {
		listener.Notify(event)
	}
}

type EventMap struct {
}

func (e *EventMap) Get(event tcell.Event) Event {
	switch ev := event.(type) {
	case *tcell.EventKey:
		return e.GetEventFromKey(ev)
	default:
		return Unknown
	}
}

func (e *EventMap) GetEventFromKey(ev *tcell.EventKey) Event {
	switch {
	case ev.Key() == tcell.KeyRune:
		switch ev.Rune() {
		case 'W', 'w':
			return MoveUp
		case 'A', 'a':
			return MoveLeft
		case 'S', 's':
			return MoveDown
		case 'D', 'd':
			return MoveRight
		case ' ':
			return PauseGame
		}
		return MoveUp
	case ev.Key() == tcell.KeyUp:
		return MoveUp
	case ev.Key() == tcell.KeyDown:
		return MoveDown
	case ev.Key() == tcell.KeyRight:
		return MoveRight
	case ev.Key() == tcell.KeyLeft:
		return MoveLeft
	case ev.Key() == tcell.KeyCtrlC:
		return ExitGame
	case ev.Key() == tcell.KeyEnter:
		return StartGame
	}
	return Unknown
}
