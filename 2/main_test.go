package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		result := part1()

		assert.Equal(t, 0, result)
	})
}

func TestPart2(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		result := part2()

		assert.Equal(t, 0, result)
	})
}
