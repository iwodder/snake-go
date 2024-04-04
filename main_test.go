package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_SnakeCanDrawIntoTheFrame(t *testing.T) {
	var dst frame
	s := snake{
		start: pos{x: 1, y: 1},
		movement: []move{
			{d: right, magnitude: 1},
		},
	}
	s.draw(&dst)

	require.Equal(t, '-', dst[1][2])
}

func Test_SnakeCanDrawDownMovementIntoTheFrame(t *testing.T) {
	var dst frame
	s := snake{
		start: pos{x: 1, y: 1},
		movement: []move{
			{d: down, magnitude: 2},
		},
	}
	s.draw(&dst)

	require.Equal(t, '|', dst[2][1])
	require.Equal(t, '|', dst[3][1])
}

func Test_SnakeCanDrawUpMovementIntoTheFrame(t *testing.T) {
	var dst frame
	s := snake{
		start: pos{x: 1, y: 1},
		movement: []move{
			{d: up, magnitude: 1},
		},
	}
	s.draw(&dst)

	require.Equal(t, '|', dst[0][1])
}

func Test_SnakeCanDrawLeftMovementIntoTheFrame(t *testing.T) {
	var dst frame
	s := snake{
		start: pos{x: 1, y: 1},
		movement: []move{
			{d: left, magnitude: 1},
		},
	}
	s.draw(&dst)

	require.Equal(t, '-', dst[1][0])
}

func Test_SnakeCanDrawMultipleMovementsIntoTheFrame(t *testing.T) {
	var dst frame
	s := snake{
		start: pos{x: 1, y: 1},
		movement: []move{
			{d: down, magnitude: 2},
			{d: right, magnitude: 2},
		},
	}
	s.draw(&dst)

	require.Equal(t, '|', dst[2][1])
	require.Equal(t, '|', dst[3][1])
	require.Equal(t, '-', dst[3][2])
	require.Equal(t, '-', dst[3][3])
}

func printFrame(f *frame) {
	for x := 0; x < len(f); x++ {
		for y := 0; y < len(f[x]); y++ {
			fmt.Printf("%c", f[x][y])
			if y < len(f[x])-1 {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}
