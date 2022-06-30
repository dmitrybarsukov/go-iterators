package extended

import (
	"testing"

	"iterator/basic"
	"iterator/util"

	"github.com/stretchr/testify/assert"
)

func TestGroupingIterator1(t *testing.T) {
	items := []int{1, 2, 3, 1, 2, 3}
	iter := GroupingIterator(basic.SliceIterator(items), util.SelfAny[int])
	assert.Equal(t, [][]int{{1, 1}, {2, 2}, {3, 3}}, util.ToSlice(iter))
}

func TestGroupingIterator2(t *testing.T) {
	items := []int{1, 2, 3, 4, 1, 2, 3}
	iter := GroupingIterator(basic.SliceIterator(items), util.SelfAny[int])
	assert.Equal(t, [][]int{{1, 1}, {2, 2}, {3, 3}, {4}}, util.ToSlice(iter))
}
