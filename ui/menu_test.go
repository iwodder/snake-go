package ui

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Menu(t *testing.T) {
	ul := Position{X: 0, Y: 0}
	setup := func() *Menu {
		ret := NewMenu(ul, 10, 10, "Test Menu")
		require.NotNil(t, ret)
		return ret
	}

	t.Run("height with no entries equals title height and border", func(t *testing.T) {
		menu := setup()
		require.Equal(t, menu.title.Height()+borderWidth*2, menu.Height())
	})

	t.Run("height equals border, title height, and all entries", func(t *testing.T) {
		menu := setup()
		menu.AddEntry("Entry 1")
		menu.AddEntry("Entry 2")

		require.Equal(t, 5, menu.Height())
	})

	t.Run("width", func(t *testing.T) {
		menu := setup()
		require.Equal(t, 10, menu.Width())
	})

	t.Run("no entries have same Y position", func(t *testing.T) {
		menu := setup()
		menu.AddEntry("Entry 1")
		menu.AddEntry("Entry 2")

		_, y := menu.title.Position()
		for _, entry := range menu.entries {
			_, obt := entry.Position()
			require.NotEqual(t, y, obt)
		}
	})

	t.Run("entries are smaller than menu", func(t *testing.T) {
		menu := setup()
		menu.AddEntry("Entry 1")
		menu.AddEntry("Entry 2")

		exp := menu.Width() - borderWidth*2
		for _, entry := range menu.entries {
			require.Equal(t, exp, entry.Width())
		}
	})
}
