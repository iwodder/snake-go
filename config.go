package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	DefaultMaxNumberOfApples      = 10
	DefaultNumberOfLives     uint = 3
	DefaultStartingLength         = 3
)

// Config holds the configuration settings for the game
type Config struct {
	maxNumberOfApples   int
	numberOfLives       uint
	snakeStartingLength int
}

// UnmarshalJSON updates the configuration using the provided JSON data.
func (c *Config) UnmarshalJSON(data []byte) error {
	type aux struct {
		MaxNumberOfApples   int  `json:"maxNumberOfApples,omitempty"`
		NumberOfLives       uint `json:"numberOfLives,omitempty"`
		SnakeStartingLength int  `json:"snakeStartingLength,omitempty"`
	}
	var a aux
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}
	c.snakeStartingLength = a.SnakeStartingLength
	c.numberOfLives = a.NumberOfLives
	c.maxNumberOfApples = a.MaxNumberOfApples
	return nil
}

// MaxNumberOfApples returns the configured maximum number of apples.
// If no value is configured, it returns the default value.
func (c *Config) MaxNumberOfApples() int {
	if c.maxNumberOfApples == 0 {
		return DefaultMaxNumberOfApples
	}
	return c.maxNumberOfApples
}

// NumberOfLives returns the configured initial number of lives.
// If no value is configured, it returns the default value.
func (c *Config) NumberOfLives() uint {
	if c.numberOfLives == 0 {
		return DefaultNumberOfLives
	}
	return c.numberOfLives
}

// SnakeStartingLength returns the configured initial length for the snake.
// If no value is configured, it returns the default value.
func (c *Config) SnakeStartingLength() int {
	if c.snakeStartingLength == 0 {
		return DefaultStartingLength
	}
	return c.snakeStartingLength
}

// LoadConfig loads the game configuration from a file.
func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open: %w", err)
	}
	var ret Config
	if err = json.NewDecoder(file).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
