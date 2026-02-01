package main

import (
	"snake/ui"
	"testing"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const moveDelta = time.Millisecond * 250

func Test_SnakeCanDrawOntoTheScreen(t *testing.T) {
	dst := setupDefaultScreen(t)

	s := snake{
		body: []cell{
			{x: 1, y: 2},
			{x: 1, y: 3},
			{x: 2, y: 3},
			{x: 3, y: 3},
			{x: 3, y: 2},
			{x: 3, y: 1},
			{x: 3, y: 0},
			{x: 2, y: 0},
			{x: 1, y: 0},
			{x: 0, y: 0},
		},
	}
	s.draw(dst)

	for _, c := range s.body {
		requireEqualContents(t, c.x, c.y, snakeRune, dst)
	}
}

func Test_Snake(t *testing.T) {
	const startingLength = 1
	initialPosition := Position{X: 5, Y: 6}

	var s *snake
	var b *ui.GameBoard

	setup := func() {
		s = newSnakeOfLength(initialPosition, startingLength)
		b = ui.NewGameBoard(Position{X: 0, Y: 0}, Position{X: 9, Y: 9})
	}

	t.Run("initial state", func(t *testing.T) {
		s = newSnake(initialPosition)

		require.Equal(t, right, s.dir)
		require.Len(t, s.body, DefaultStartingLength)
		require.Equal(t, initialPosition, s.headPos())
	})

	t.Run("responds to events", func(t *testing.T) {
		t.Run("down event", func(t *testing.T) {
			setup()

			simulate(s, b, MoveDown)

			exp := Position{X: initialPosition.X, Y: initialPosition.Y + 1}
			require.Equal(t, exp, s.headPos())
		})

		t.Run("up event", func(t *testing.T) {
			setup()

			simulate(s, b, MoveUp)

			exp := Position{X: initialPosition.X, Y: initialPosition.Y - 1}
			require.Equal(t, exp, s.headPos())
		})

		t.Run("left event", func(t *testing.T) {
			setup()

			simulate(s, b, MoveDown, MoveLeft)

			exp := Position{X: initialPosition.X - 1, Y: initialPosition.Y + 1}
			require.Equal(t, exp, s.headPos())
		})

		t.Run("right event", func(t *testing.T) {
			setup()

			simulate(s, b, MoveDown, MoveRight)

			exp := Position{X: initialPosition.X + 1, Y: initialPosition.Y + 1}
			require.Equal(t, exp, s.headPos())
		})
	})

	t.Run("can change direction", func(t *testing.T) {
		t.Run("left", func(t *testing.T) {
			setup()
			s.changeDirection(up)
			// snake starts faced right

			s.changeDirection(left)
			require.Equal(t, left, s.dir)
		})

		t.Run("right", func(t *testing.T) {
			setup()
			s.changeDirection(up)
			// snake starts faced right

			s.changeDirection(right)
			require.Equal(t, right, s.dir)
		})

		t.Run("up", func(t *testing.T) {
			setup()
			s.changeDirection(up)
			require.Equal(t, up, s.dir)
		})

		t.Run("down", func(t *testing.T) {
			setup()
			s.changeDirection(down)
			require.Equal(t, down, s.dir)
		})
	})

	t.Run("doesn't double back on itself", func(t *testing.T) {
		t.Run("right to left", func(t *testing.T) {
			setup()

			s.changeDirection(left)

			require.Equal(t, right, s.dir)
		})

		t.Run("left to right", func(t *testing.T) {
			setup()
			s.dir = left // snake starts faced right

			s.changeDirection(right)

			require.Equal(t, left, s.dir)
		})

		t.Run("up to down", func(t *testing.T) {
			setup()
			s.dir = up

			s.changeDirection(down)

			require.Equal(t, up, s.dir)
		})

		t.Run("down to up", func(t *testing.T) {
			setup()
			s.dir = down

			s.changeDirection(up)

			require.Equal(t, down, s.dir)
		})
	})

	t.Run("movement", func(t *testing.T) {
		t.Run("up", func(t *testing.T) {
			setup()

			simulate(s, b, MoveUp)

			exp := cell{x: initialPosition.X, y: initialPosition.Y - 1}
			require.Equal(t, up, s.dir)
			require.Equal(t, exp, s.body[0])
		})

		t.Run("down", func(t *testing.T) {
			setup()

			simulate(s, b, MoveDown)

			exp := cell{x: initialPosition.X, y: initialPosition.Y + 1}
			require.Equal(t, down, s.dir)
			require.Equal(t, exp, s.body[0])
		})

		t.Run("right", func(t *testing.T) {
			setup()

			simulate(s, b, MoveUp)
			simulate(s, b, MoveRight)

			exp := cell{x: initialPosition.X + 1, y: initialPosition.Y - 1}
			require.Equal(t, right, s.dir)
			require.Equal(t, exp, s.body[0])
		})

		t.Run("left", func(t *testing.T) {
			setup()

			simulate(s, b, MoveUp)
			simulate(s, b, MoveLeft)

			exp := cell{x: initialPosition.X - 1, y: initialPosition.Y - 1}
			require.Equal(t, left, s.dir)
			require.Equal(t, exp, s.body[0])
		})

		t.Run("stays on board (right)", func(t *testing.T) {
			setup()

			simulate(s, b, MoveRight, MoveRight, MoveRight, MoveRight, MoveRight)

			require.Equal(t, Position{X: b.Right(), Y: initialPosition.Y}, s.headPos())
		})

		t.Run("stays on board (up)", func(t *testing.T) {
			setup()

			simulate(s, b, MoveUp, MoveUp, MoveUp, MoveUp, MoveUp)

			require.Equal(t, Position{X: initialPosition.X, Y: b.Top()}, s.headPos())
		})

		t.Run("stays on board (down)", func(t *testing.T) {
			setup()

			simulate(s, b, MoveDown, MoveDown, MoveDown, MoveDown, MoveDown)

			require.Equal(t, Position{X: initialPosition.X, Y: b.Bottom()}, s.headPos())
		})

		t.Run("stays on board (left)", func(t *testing.T) {
			setup()

			simulate(s, b, MoveUp, MoveLeft, MoveLeft, MoveLeft, MoveLeft, MoveLeft)

			require.Equal(t, Position{X: b.Left(), Y: initialPosition.Y - 1}, s.headPos())
		})

		t.Run("doesn't move until direction is away from edge (right)", func(t *testing.T) {
			setup()

			simulate(s, b, MoveRight, MoveRight, MoveRight)
			require.Equal(t, Position{X: b.Right(), Y: initialPosition.Y}, s.headPos(),
				"snake not to right edge")

			simulate(s, b, MoveRight)
			require.Equal(t, Position{X: b.Right(), Y: initialPosition.Y}, s.headPos(),
				"snake moved after hitting right edge")

			simulate(s, b, MoveUp)
			require.Equal(t, Position{X: b.Right(), Y: initialPosition.Y - 1}, s.headPos())
		})

		t.Run("doesn't move until direction is away from edge (left)", func(t *testing.T) {
			setup()

			simulate(s, b, MoveUp, MoveLeft, MoveLeft, MoveLeft, MoveLeft, MoveLeft)
			require.Equal(t, Position{X: b.Left(), Y: initialPosition.Y - 1}, s.headPos(),
				"snake not to left edge")

			simulate(s, b, MoveLeft)
			require.Equal(t, Position{X: b.Left(), Y: initialPosition.Y - 1}, s.headPos(),
				"snake moved after hitting left edge")

			simulate(s, b, MoveDown)
			require.Equal(t, Position{X: b.Left(), Y: initialPosition.Y}, s.headPos())
		})

		t.Run("doesn't move until direction is away from edge (top)", func(t *testing.T) {
			setup()

			simulate(s, b, MoveUp, MoveUp, MoveUp, MoveUp)
			require.Equal(t, Position{X: initialPosition.X, Y: b.Top()}, s.headPos(),
				"snake not to top edge")

			simulate(s, b, MoveUp)
			require.Equal(t, Position{X: initialPosition.X, Y: b.Top()}, s.headPos(),
				"snake moved after hitting top edge")

			simulate(s, b, MoveLeft)
			require.Equal(t, Position{X: initialPosition.X - 1, Y: b.Top()}, s.headPos())
		})

		t.Run("doesn't move until direction is away from edge (bottom)", func(t *testing.T) {
			setup()

			simulate(s, b, MoveDown, MoveDown, MoveDown, MoveDown)
			require.Equal(t, Position{X: initialPosition.X, Y: b.Bottom()}, s.headPos(),
				"snake not to bottom edge")

			simulate(s, b, MoveDown)
			require.Equal(t, Position{X: initialPosition.X, Y: b.Bottom()}, s.headPos(),
				"snake moved after hitting bottom edge")

			simulate(s, b, MoveRight)
			require.Equal(t, Position{X: initialPosition.X + 1, Y: b.Bottom()}, s.headPos())
		})

		t.Run("is 4 cells per second", func(t *testing.T) {
			rate := time.Second / 20
			ticker := time.NewTicker(rate)
			startPos := Position{X: 0, Y: 5}

			setup()
			s.body[0].x = startPos.X
			s.body[0].y = startPos.Y

			ticks := 0
			for range ticker.C {
				s.move(b, rate)
				if ticks += 1; ticks == 20 {
					ticker.Stop()
					break
				}
			}
			exp := startPos
			exp.X += 4
			require.Equal(t, exp, s.headPos())
		})
	})

	t.Run("grows by eating apples", func(t *testing.T) {
		setup()

		as := apples{
			{pos: initialPosition, eaten: false},
			{pos: Position{X: 1, Y: 1}, eaten: false},
		}
		s.eat(as)

		require.Len(t, s.body, 2)
		require.True(t, as[0].eaten)
		require.False(t, as[1].eaten)
	})

	t.Run("snake moving in left circle crashes", func(t *testing.T) {
		setup()
		s.body = []cell{
			{x: 3, y: 3},
			{x: 2, y: 3},
			{x: 2, y: 2},
			{x: 3, y: 2},
			{x: 3, y: 3},
		}

		require.True(t, s.crashed())
	})

	t.Run("snake moving in right circle crashes", func(t *testing.T) {
		setup()
		s.body = []cell{
			{x: 3, y: 3},
			{x: 4, y: 3},
			{x: 4, y: 2},
			{x: 3, y: 2},
			{x: 3, y: 3},
		}

		require.True(t, s.crashed())
	})

	t.Run("moving in straight line (right) doesn't crash", func(t *testing.T) {
		setup()
		s.body = []cell{
			{x: 3, y: 3},
			{x: 4, y: 3},
			{x: 5, y: 3},
		}

		require.False(t, s.crashed())
	})

	t.Run("moving in straight line (left) doesn't crash", func(t *testing.T) {
		setup()
		s.body = []cell{
			{x: 3, y: 3},
			{x: 2, y: 3},
			{x: 1, y: 3},
		}

		require.False(t, s.crashed())
	})

	t.Run("moving in straight line (up) doesn't crash", func(t *testing.T) {
		setup()
		s.body = []cell{
			{x: 3, y: 3},
			{x: 3, y: 2},
			{x: 3, y: 1},
		}

		require.False(t, s.crashed())
	})

	t.Run("moving in straight line (down) doesn't crash", func(t *testing.T) {
		setup()
		s.body = []cell{
			{x: 3, y: 3},
			{x: 3, y: 4},
			{x: 3, y: 5},
		}

		require.False(t, s.crashed())
	})

	t.Run("reset restores initial state", func(t *testing.T) {
		setup()
		s.body = []cell{
			{x: 3, y: 3},
			{x: 3, y: 4},
			{x: 3, y: 5},
		}
		simulate(s, b, MoveUp, MoveUp, MoveRight, MoveLeft)

		pos := Position{1, 1}
		s.ResetTo(pos)

		require.Equal(t, right, s.dir)
		require.Len(t, s.body, 1)
		require.Equal(t, pos, s.headPos())
		require.Equal(t, s.startingLength, s.Length())
	})

	t.Run("speed increases by 25% when doubling in length", func(t *testing.T) {
		setup()
		as := apples{
			{pos: initialPosition, eaten: false},
			{pos: Position{X: 1, Y: 1}, eaten: false},
		}

		speed := s.moveDelay

		for range startingLength * 2 {
			s.eat(as)
		}

		assert.Equal(t, float64(speed)*0.75, float64(s.moveDelay))
		assert.Equal(t, startingLength*2, s.lastLength)
	})
}

func requireEqualContents(t *testing.T, x, y int, exp rune, scn tcell.SimulationScreen) {
	act, _, _, _ := scn.GetContent(x, y)
	require.EqualValues(t, exp, act, "position (x=%d,Y=%d) expected '%c', but was '%c'", x, y, exp, act)
}

func setupDefaultScreen(t *testing.T) tcell.SimulationScreen {
	return setupScreen(t, 80, 80)
}

func simulate(s *snake, b *ui.GameBoard, events ...Event) {
	for _, event := range events {
		s.Notify(event)
		s.move(b, moveDelta)
	}
}

func setupScreen(t *testing.T, height, width int) tcell.SimulationScreen {
	ret := tcell.NewSimulationScreen("")
	require.NoError(t, ret.Init())
	ret.SetSize(height, width)
	return ret
}
