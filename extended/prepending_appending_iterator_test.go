package extended

import (
	"testing"

	"iterator/basic"
	"iterator/util"

	"github.com/stretchr/testify/assert"
)

func TestAppendingIteratorSeveralItems(t *testing.T) {
	iter := AppendingIterator(basic.SliceIterator([]int{1, 2, 3}), 4, 5, 6)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, util.ToSlice(iter))
}

func TestAppendingIteratorNoItems(t *testing.T) {
	iter := AppendingIterator(basic.SliceIterator([]int{1, 2, 3}))
	assert.Equal(t, []int{1, 2, 3}, util.ToSlice(iter))
}

func TestPrependingIteratorSeveralItems(t *testing.T) {
	iter := PrependingIterator(basic.SliceIterator([]int{1, 2, 3}), 4, 5, 6)
	assert.Equal(t, []int{4, 5, 6, 1, 2, 3}, util.ToSlice(iter))
}

func TestPrependingIteratorNoItems(t *testing.T) {
	iter := PrependingIterator(basic.SliceIterator([]int{1, 2, 3}))
	assert.Equal(t, []int{1, 2, 3}, util.ToSlice(iter))
}

func TestPrependingAppendingIteratorChain(t *testing.T) {
	iter := PrependingIterator(
		AppendingIterator(
			PrependingIterator(
				AppendingIterator(
					basic.SliceIterator([]int{1}),
					4, 5,
				), 6, 7,
			),
			8,
		), 9,
	)
	assert.Equal(t, []int{9, 6, 7, 1, 4, 5, 8}, util.ToSlice(iter))
	assert.Equal(t, []int{4, 5, 8}, iter.(*prependingAppendingIterator[int]).itemsToAppend)
	assert.Equal(t, []int{9, 6, 7}, iter.(*prependingAppendingIterator[int]).itemsToPrepend)
}
