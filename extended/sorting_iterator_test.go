package extended

import (
	"testing"

	"iterator/basic"
	"iterator/util"

	"github.com/stretchr/testify/assert"
)

func TestSortingIterator(t *testing.T) {
	iter := SortingIterator(basic.SliceIterator([]int{5, 1, 4, 2, 3}), func(i, j int) bool { return i < j })
	assert.Equal(t, []int{1, 2, 3, 4, 5}, util.ToSlice(iter))
}

func TestSortingIteratorAscending(t *testing.T) {
	iter := SortingIteratorAsc(basic.SliceIterator([]int{5, 1, 4, 2, 3}), util.Self[int])
	assert.Equal(t, []int{1, 2, 3, 4, 5}, util.ToSlice(iter))
}

func TestSortingIteratorDescending(t *testing.T) {
	iter := SortingIteratorDesc(basic.SliceIterator([]int{5, 1, 4, 2, 3}), util.Self[int])
	assert.Equal(t, []int{5, 4, 3, 2, 1}, util.ToSlice(iter))
}
