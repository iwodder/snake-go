package ui

import "testing"

func Test_EntityRendering(t *testing.T) {
	t.Run("snake", func(t *testing.T) {
		dst := setup(t)

		s := SnakeRenderer{
			Body: []Position{
				{X: 1, Y: 2},
				{X: 1, Y: 3},
				{X: 2, Y: 3},
				{X: 3, Y: 3},
				{X: 3, Y: 2},
				{X: 3, Y: 1},
				{X: 3, Y: 0},
				{X: 2, Y: 0},
				{X: 1, Y: 0},
				{X: 0, Y: 0},
			},
		}
		s.Draw(dst)

		for _, c := range s.Body {
			requireEqualContents(t, c.X, c.Y, snakeRune, dst)
		}
	})

	t.Run("apple", func(t *testing.T) {
		scn := setup(t)

		ar := AppleRenderer{
			Pos: Position{X: 1, Y: 1},
		}

		ar.Draw(scn)

		requireEqualContents(t, 1, 1, appleRune, scn)
	})
}
