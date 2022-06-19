package extended

import (
	"testing"

	"iterator/basic"

	"github.com/stretchr/testify/assert"
)

func TestFilteringIteratorNone(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	iter := FilteringIterator(basic.Slice(items), func(i int) bool { return false })
	result := make([]int, 0)
	for iter.HasNext() {
		result = append(result, iter.Next())
	}
	assert.Empty(t, result)
}

func TestFilteringIteratorAll(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	iter := FilteringIterator(basic.Slice(items), func(i int) bool { return true })
	result := make([]int, 0)
	for iter.HasNext() {
		result = append(result, iter.Next())
	}
	assert.Equal(t, items, result)
}

func TestFilteringIteratorSome(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	iter := FilteringIterator(basic.Slice(items), func(i int) bool { return i%2 == 0 })
	result := make([]int, 0)
	for iter.HasNext() {
		result = append(result, iter.Next())
	}
	assert.Equal(t, []int{2, 4}, result)
}
