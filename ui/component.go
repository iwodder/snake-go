package ui

import (
	"slices"

	"github.com/gdamore/tcell/v2"
)

type Component interface {
	Draw(tcell.Screen)
	Add(c Component) error
	Remove(c Component) error
}

type leaf struct{}

func (l leaf) Add(Component) error {
	return nil
}

func (l leaf) Remove(Component) error {
	return nil
}

type composite []Component

func (c *composite) Draw(scrn tcell.Screen) {
	for _, comp := range *c {
		comp.Draw(scrn)
	}
}

func (c *composite) Add(comp Component) error {
	*c = append(*c, comp)
	return nil
}

func (c *composite) Remove(comp Component) error {
	*c = slices.DeleteFunc(*c, func(other Component) bool {
		return comp == other
	})
	return nil
}
