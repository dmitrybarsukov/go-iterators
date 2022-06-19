package extended

import (
	"testing"

	"iterator/basic"

	"github.com/stretchr/testify/assert"
)

func TestSkippingIterator(t *testing.T) {
	items := []int{1, 2, 3, 4}
	iter := SkippingIterator(basic.SliceIterator(items), 2)
	result := make([]int, 0)
	for iter.HasNext() {
		result = append(result, iter.Next())
	}
	assert.Equal(t, []int{3, 4}, result)
}

func TestSkippingIteratorSkipAll(t *testing.T) {
	items := []int{1, 2, 3, 4}
	iter := SkippingIterator(basic.SliceIterator(items), 20)
	result := make([]int, 0)
	for iter.HasNext() {
		result = append(result, iter.Next())
	}
	assert.Equal(t, []int{}, result)
}
