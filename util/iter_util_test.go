package util

import (
	"testing"

	"iterator/basic"

	"github.com/stretchr/testify/assert"
)

func TestToSlice(t *testing.T) {
	items := []int{1, 2, 3, 4}
	iter := basic.Slice(items)
	assert.Equal(t, items, ToSlice(iter))
}

func TestToMap(t *testing.T) {
	items := []int{1, 2, 3, 4}
	iter := basic.Slice(items)
	actual := ToMap(iter, func(i int) int { return i - 1 })
	expected := map[int]int{0: 1, 1: 2, 2: 3, 3: 4}
	assert.Equal(t, expected, actual)
}

func TestToMapWithValue(t *testing.T) {
	items := []int{1, 2, 3, 4}
	iter := basic.Slice(items)
	actual := ToMapWithValue(iter, func(i int) int { return i - 1 }, func(i int) int { return i + 1 })
	expected := map[int]int{0: 2, 1: 3, 2: 4, 3: 5}
	assert.Equal(t, expected, actual)
}

func TestCount(t *testing.T) {
	items := []int{1, 2, 3, 4}
	iter := basic.Slice(items)
	assert.Equal(t, 4, Count(iter))
}

func TestCountMatching(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	iter := basic.Slice(items)
	assert.Equal(t, 3, CountMatching(iter, func(i int) bool { return i%2 == 1 }))
}

func TestFirst(t *testing.T) {
	_, ok := First(basic.Slice([]int{}))
	assert.False(t, ok)

	value, ok := First(basic.Slice([]int{1, 2, 3}))
	assert.True(t, ok)
	assert.Equal(t, 1, value)
}

func TestFirstOrZeroValue(t *testing.T) {
	value := FirstOrZeroValue(basic.Slice([]int{}))
	assert.Zero(t, value)
}

func TestFirstOrDefault(t *testing.T) {
	value := FirstOrDefault(basic.Slice([]int{}), 5)
	assert.Equal(t, 5, value)
}

func TestLast(t *testing.T) {
	_, ok := Last(basic.Slice([]int{}))
	assert.False(t, ok)

	value, ok := Last(basic.Slice([]int{1, 2, 3}))
	assert.True(t, ok)
	assert.Equal(t, 3, value)
}

func TestLastOrZeroValue(t *testing.T) {
	value := LastOrZeroValue(basic.Slice([]int{}))
	assert.Zero(t, value)
}

func TestLastOrDefault(t *testing.T) {
	value := LastOrDefault(basic.Slice([]int{}), 5)
	assert.Equal(t, 5, value)
}
