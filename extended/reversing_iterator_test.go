package extended

import (
	"testing"

	"iterator/basic"
	"iterator/util"

	"github.com/stretchr/testify/assert"
)

func TestReversingIterator(t *testing.T) {
	iter := ReversingIterator(basic.SliceIterator([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, []int{5, 4, 3, 2, 1}, util.ToSlice(iter))
}
