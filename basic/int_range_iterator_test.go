package basic

import (
	"testing"

	"iterator/commons"
	"iterator/util"

	"github.com/stretchr/testify/assert"
)

func TestIntRangeIteratorForward(t *testing.T) {
	iter := IntRangeIterator(1, 9, 2)
	assert.Equal(t, []int{1, 3, 5, 7}, util.ToSlice(iter))
}

func TestIntRangeIteratorBackward(t *testing.T) {
	iter := IntRangeIterator(9, 1, -2)
	assert.Equal(t, []int{9, 7, 5, 3}, util.ToSlice(iter))
}

func TestIntRangeIteratorWrongRanges(t *testing.T) {
	testCases := []struct {
		name   string
		action func()
		err    error
	}{
		{
			name:   "zero step",
			action: func() { IntRangeIterator(1, 2, 0) },
			err:    commons.ErrArgumentStepIsZero,
		},
		{
			name:   "range positive, step negative",
			action: func() { IntRangeIterator(1, 2, -1) },
			err:    commons.ErrArgumentStepHasWrongSign,
		},
		{
			name:   "range negative, step positive",
			action: func() { IntRangeIterator(2, 1, 1) },
			err:    commons.ErrArgumentStepHasWrongSign,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.PanicsWithValue(t, testCase.err, func() {
				testCase.action()
			})
		})
	}
}
