package extended

import (
	"testing"

	"iterator/basic"
	"iterator/commons"

	"github.com/stretchr/testify/assert"
)

func TestLimitingIterator(t *testing.T) {
	items := []int{1, 2, 3, 4}
	iter := LimitingIterator(basic.Slice(items), 2)
	result := make([]int, 0)
	for iter.HasNext() {
		result = append(result, iter.Next())
	}
	assert.Equal(t, []int{1, 2}, result)
}

func TestLimitingIteratorUnlimited(t *testing.T) {
	items := []int{1, 2, 3, 4}
	iter := LimitingIterator(basic.Slice(items), 20)
	result := make([]int, 0)
	for iter.HasNext() {
		result = append(result, iter.Next())
	}
	assert.Equal(t, items, result)
}

func TestLimitedIteratorPanicsIfIteratorEnded(t *testing.T) {
	items := []int{1, 2}
	iter := LimitingIterator(basic.Slice(items), 1)
	for iter.HasNext() {
		_ = iter.Next()
	}
	err := func() (err error) {
		defer func() {
			if errRec, ok := recover().(error); ok {
				err = errRec
			}
		}()
		iter.Next()
		return
	}()

	assert.Equal(t, commons.ErrIterEnded, err)
}
