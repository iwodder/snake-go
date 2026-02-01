package ui

import (
	"testing"

	"github.com/gdamore/tcell/v2"
)

func Test_Composite(t *testing.T) {
	var spy *keyEventSpy
	var comp *composite

	setup := func() {
		spy = &keyEventSpy{t: t}
		comp = &composite{}
		if err := comp.Add(spy); err != nil {
			t.Fatalf("adding component failed: %v", err)
		}
	}

	t.Run("propagates key event to children", func(t *testing.T) {
		setup()

		comp.Handle(&tcell.EventKey{})

		spy.AssertHandleKeyEventWasCalled()
	})

	t.Run("can remove child component", func(t *testing.T) {
		setup()

		if err := comp.Remove(spy); err != nil {
			t.Fatalf("removing component failed: %v", err)
		}
		if exp, act := len(comp.components), 0; exp != act {
			t.Fatalf("expected %d components, got %d", exp, act)
		}
	})

	t.Run("notifies self of key event", func(t *testing.T) {
		wasRun := new(bool)

		comp.SetKeyEventCallback(func(*tcell.EventKey) {
			*wasRun = true
		})

		comp.Handle(&tcell.EventKey{})

		if !*wasRun {
			t.Fatal("callback was not called")
		}
	})
}

type keyEventSpy struct {
	t   *testing.T
	key *tcell.EventKey
	leaf
}

func (k *keyEventSpy) Draw(tcell.Screen) {}

func (k *keyEventSpy) handleKeyEvent(key *tcell.EventKey) {
	k.key = key
}

func (k *keyEventSpy) SetKeyEventCallback(func(*tcell.EventKey)) {
}

func (k *keyEventSpy) AssertHandleKeyEventWasCalled() {
	if k.key == nil {
		k.t.Fatal("handleKeyEvent was not called")
	}
}
