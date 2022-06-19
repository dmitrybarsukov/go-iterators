package basic

import (
	"testing"

	"iterator/commons"
	"iterator/util"

	"github.com/stretchr/testify/assert"
)

func TestSliceIteratorEmpty(t *testing.T) {
	items := make([]int, 0)
	iter := SliceIterator(items)
	assert.Zero(t, 0, util.Count(iter))
}

func TestSliceIteratorMany(t *testing.T) {
	items := []int{1, 2, 5, 7, 0}
	iter := SliceIterator(items)
	assert.Equal(t, items, util.ToSlice(iter))
}

func TestSliceIteratorPanicsIfIteratorEnded(t *testing.T) {
	items := []int{1, 2}
	iter := SliceIterator(items)
	for iter.HasNext() {
		_ = iter.Next()
	}
	assert.PanicsWithValue(t, commons.ErrIterEnded, func() {
		iter.Next()
	})
}

func TestSliceIteratorNil(t *testing.T) {
	var items []int
	iter := SliceIterator(items)
	assert.False(t, iter.HasNext())
}
