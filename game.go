package main

import "github.com/gdamore/tcell/v2"

type keyListener interface {
	notify(event *tcell.EventKey)
}

type keyListeners []keyListener

func (k keyListeners) post(event *tcell.EventKey) {
	for _, listener := range k {
		listener.notify(event)
	}
}

type game struct {
	kl  keyListeners
	scn tcell.Screen
}

func (g *game) pollEvents() {
	for {
		switch ev := g.scn.PollEvent().(type) {
		case *tcell.EventKey:
			g.kl.post(ev)
		case nil: // screen finalized
			return
		}
	}
}

func (g *game) registerKeyListener(kl keyListener) {
	g.kl = append(g.kl, kl)
}
