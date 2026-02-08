package ui

import (
	"testing"

	"github.com/gdamore/tcell/v2"
)

func Test_DrawBorder(t *testing.T) {
	setup := func() tcell.SimulationScreen {
		scrn := tcell.NewSimulationScreen("utf-8")
		if err := scrn.Init(); err != nil {
			t.Fatalf("failed to initialize screen: %v", err)
		}
		t.Cleanup(scrn.Fini)
		return scrn
	}

	t.Run("panics on nil screen", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic, but none occurred")
			}
		}()
		drawBorder(Position{}, 0, 0, tcell.Style{}, nil)
	})

	t.Run("draws border", func(t *testing.T) {
		scrn := setup()

		drawBorder(Position{X: 0, Y: 0}, 5, 5, tcell.StyleDefault, scrn)

		assertEqualContents(t, Position{X: 0, Y: 0}, tcell.RuneULCorner, scrn)
		assertEqualContents(t, Position{X: 4, Y: 0}, tcell.RuneURCorner, scrn)
		assertEqualContents(t, Position{X: 0, Y: 4}, tcell.RuneLLCorner, scrn)
		assertEqualContents(t, Position{X: 4, Y: 4}, tcell.RuneLRCorner, scrn)
	})

	t.Run("panics if width is less than one", func(t *testing.T) {
		scrn := setup()

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic, but none occurred")
			}
		}()
		drawBorder(Position{}, 0, 1, tcell.StyleDefault, scrn)
	})

	t.Run("panics if height is less than one", func(t *testing.T) {
		scrn := setup()

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic, but none occurred")
			}
		}()
		drawBorder(Position{}, 1, 0, tcell.StyleDefault, scrn)
	})
}

func assertEqualContents(t *testing.T, pos Position, exp rune, scrn tcell.SimulationScreen) {
	act, _, _, _ := scrn.GetContent(pos.X, pos.Y)
	if act != exp {
		t.Errorf("expected %q for %#v, got %q.", exp, pos, act)
	}
}
