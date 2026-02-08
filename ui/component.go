package ui

import (
	"slices"

	"github.com/gdamore/tcell/v2"
)

type Component interface {
	Draw(tcell.Screen)
	Add(c Component) error
	Remove(c Component) error
	Width() int
	Height() int
}

type keyHandler interface {
	handleKeyEvent(*tcell.EventKey)
	SetKeyEventCallback(callback func(*tcell.EventKey))
}

type EventHandler interface {
	Handle(tcell.Event)
}

type leaf struct{}

func (l leaf) Add(Component) error {
	return nil
}

func (l leaf) Remove(Component) error {
	return nil
}

type composite struct {
	components       []Component
	keyEventCallback func(*tcell.EventKey)
}

func (c *composite) Draw(scrn tcell.Screen) {
	for _, comp := range c.components {
		comp.Draw(scrn)
	}
}

func (c *composite) Handle(event tcell.Event) {
	switch ev := event.(type) {
	case *tcell.EventKey:
		c.handleKeyEvent(ev)
		for _, comp := range c.components {
			if keyHandler, ok := comp.(keyHandler); ok {
				keyHandler.handleKeyEvent(ev)
			}
		}
	}
}

func (c *composite) Add(comp Component) error {
	c.components = append(c.components, comp)
	return nil
}

func (c *composite) Remove(comp Component) error {
	c.components = slices.DeleteFunc(c.components, func(other Component) bool {
		return comp == other
	})
	return nil
}

func (c *composite) SetKeyEventCallback(callback func(*tcell.EventKey)) {
	c.keyEventCallback = callback
}

func (c *composite) handleKeyEvent(ev *tcell.EventKey) {
	if c.keyEventCallback != nil {
		c.keyEventCallback(ev)
	}
}
