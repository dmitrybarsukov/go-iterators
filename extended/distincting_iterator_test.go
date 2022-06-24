package extended

import (
	"testing"

	"iterator/basic"
	"iterator/util"

	"github.com/stretchr/testify/assert"
)

func TestDistinctingIteratorEmpty(t *testing.T) {
	iter := DistinctingIterator(basic.SliceIterator([]int{}), util.SelfAny[int])
	assert.Empty(t, util.ToSlice(iter))
}

func TestDistinctingIteratorAllUnique(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	iter := DistinctingIterator(basic.SliceIterator(items), util.SelfAny[int])
	result := util.ToSlice(iter)
	assert.Equal(t, items, result)
}

func TestDistinctingFilterUnique(t *testing.T) {
	items := []int{1, 2, 3, 1, 5, 4, 3, 1, 2, 7, 7, 7, 1, 2, 3}
	iter := DistinctingIterator(basic.SliceIterator(items), util.SelfAny[int])
	result := util.ToSlice(iter)
	assert.Equal(t, []int{1, 2, 3, 5, 4, 7}, result)
}
