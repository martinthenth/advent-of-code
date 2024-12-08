package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		assert.Equal(t, true, isSafe([]int{7, 6, 4, 2, 1}))
		assert.Equal(t, false, isSafe([]int{1, 2, 7, 8, 9}))
		assert.Equal(t, false, isSafe([]int{9, 7, 6, 2, 1}))
		assert.Equal(t, false, isSafe([]int{1, 3, 2, 4, 5}))
		assert.Equal(t, false, isSafe([]int{8, 6, 4, 4, 1}))
		assert.Equal(t, true, isSafe([]int{1, 3, 6, 7, 9}))
	})
}

func TestPart2(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		assert.Equal(t, true, isSafeWithTolerance([]int{7, 6, 4, 2, 1}))
		assert.Equal(t, false, isSafeWithTolerance([]int{1, 2, 7, 8, 9}))
		assert.Equal(t, false, isSafeWithTolerance([]int{9, 7, 6, 2, 1}))
		assert.Equal(t, true, isSafeWithTolerance([]int{1, 3, 2, 4, 5}))
		assert.Equal(t, true, isSafeWithTolerance([]int{8, 6, 4, 4, 1}))
		assert.Equal(t, true, isSafeWithTolerance([]int{1, 3, 6, 7, 9}))
	})
}
