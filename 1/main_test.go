package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		numbers1 := []int{3, 4, 2, 1, 3, 3}
		numbers2 := []int{4, 3, 5, 3, 9, 3}

		result := part1(numbers1, numbers2)

		assert.Equal(t, 11, result)
	})
}

func TestPart2(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		numbers1 := []int{3, 4, 2, 1, 3, 3}
		numbers2 := []int{4, 3, 5, 3, 9, 3}

		result := part2(numbers1, numbers2)

		assert.Equal(t, 31, result)
	})
}
