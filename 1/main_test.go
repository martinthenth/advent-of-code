package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		result := part1("example")

		assert.Equal(t, result, 1)
	})
}

func TestPart2(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		result := part2("example")

		assert.Equal(t, result, 0)
	})
}
