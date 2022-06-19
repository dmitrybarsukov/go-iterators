package extended

import (
	"testing"

	"iterator/basic"
	"iterator/util"

	"github.com/stretchr/testify/assert"
)

func TestSortingIteratorAscending(t *testing.T) {
	iter := SortingIterator(basic.SliceIterator([]int{5, 1, 4, 2, 3}), func(i int) int { return i }, false)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, util.ToSlice(iter))
}

func TestSortingIteratorDescending(t *testing.T) {
	iter := SortingIterator(basic.SliceIterator([]int{5, 1, 4, 2, 3}), func(i int) int { return i }, true)
	assert.Equal(t, []int{5, 4, 3, 2, 1}, util.ToSlice(iter))
}
