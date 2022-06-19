package basic

import (
	"testing"

	"iterator/commons"
	"iterator/util"

	"github.com/stretchr/testify/assert"
)

func TestRepeatIterator(t *testing.T) {
	iter := RepeatIterator(1, 5)
	assert.Equal(t, []int{1, 1, 1, 1, 1}, util.ToSlice(iter))
}

func TestRepeatPanic(t *testing.T) {
	assert.PanicsWithValue(t, commons.ErrArgumentCountNegative, func() {
		RepeatIterator(1, -1)
	})
}
