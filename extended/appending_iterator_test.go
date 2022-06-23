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

func TestAppendingIteratorChain(t *testing.T) {
	iter := AppendingIterator(
		AppendingIterator(
			AppendingIterator(
				basic.SliceIterator([]int{1, 2, 3}),
				4),
			5),
		6,
	)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, util.ToSlice(iter))
	assert.Equal(t, []int{4, 5, 6}, iter.(*appendingIterator[int]).itemsToAppend)
}
