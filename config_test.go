package main

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const exampleConfig = "" +
	`
{
	"maxNumberOfApples": 15,
	"numberOfLives": 3,
	"snakeStartingLength": 5
}
`

var expectedConfig = Config{
	maxNumberOfApples:   15,
	numberOfLives:       3,
	snakeStartingLength: 5,
}

func Test_Config(t *testing.T) {
	t.Run("successfully creates config from json", func(t *testing.T) {
		var cfg Config
		dec := json.NewDecoder(strings.NewReader(exampleConfig))

		require.NoError(t, dec.Decode(&cfg))
		require.Equal(t, expectedConfig, cfg)
	})

	t.Run("returns default max number of apples when not defined", func(t *testing.T) {
		updatedConfig := strings.Replace(exampleConfig, "\"maxNumberOfApples\": 15,", "", -1)

		var cfg Config
		dec := json.NewDecoder(strings.NewReader(updatedConfig))
		require.NoError(t, dec.Decode(&cfg))

		require.Equal(t, DefaultMaxNumberOfApples, cfg.MaxNumberOfApples())
	})

	t.Run("returns default number of lives when not defined", func(t *testing.T) {
		updatedConfig := strings.Replace(exampleConfig, "\"numberOfLives\": 3,", "", -1)

		var cfg Config
		dec := json.NewDecoder(strings.NewReader(updatedConfig))
		require.NoError(t, dec.Decode(&cfg))

		require.Equal(t, DefaultNumberOfLives, cfg.NumberOfLives())
	})

	t.Run("returns default starting length when not defined", func(t *testing.T) {
		updatedConfig := strings.Replace(exampleConfig, "\"snakeStartingLength\": 5", "", -1)
		updatedConfig = strings.Replace(updatedConfig, "3,", "3", -1)

		var cfg Config
		dec := json.NewDecoder(strings.NewReader(updatedConfig))
		require.NoError(t, dec.Decode(&cfg))

		require.Equal(t, DefaultStartingLength, cfg.SnakeStartingLength())
	})
}

func Test_LoadConfigFromFile(t *testing.T) {
	dir := t.TempDir()
	file, err := os.CreateTemp(dir, "*.json")
	require.NoError(t, err, "unable to create temp file")
	_, err = file.ReadFrom(strings.NewReader(exampleConfig))
	require.NoError(t, err, "unable to write to file")

	act, err := LoadConfig(file.Name())

	require.NoError(t, err)
	require.Equal(t, &expectedConfig, act)
}
