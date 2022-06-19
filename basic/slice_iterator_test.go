package basic

import (
	"testing"

	"iterator/commons"

	"github.com/stretchr/testify/assert"
)

func TestSliceIteratorEmpty(t *testing.T) {
	items := make([]int, 0)
	iter := Slice(items)
	count := 0
	for iter.HasNext() {
		_ = iter.Next()
		count += 1
	}
	assert.Equal(t, 0, count)
}

func TestSliceIteratorMany(t *testing.T) {
	items := []int{1, 2, 5, 7, 0}
	iter := Slice(items)
	result := make([]int, 0)
	for iter.HasNext() {
		result = append(result, iter.Next())
	}
	assert.Equal(t, items, result)
}

func TestSliceIteratorPanicsIfIteratorEnded(t *testing.T) {
	items := []int{1, 2}
	iter := Slice(items)
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

func TestSliceIteratorNil(t *testing.T) {
	var items []int
	iter := Slice(items)
	assert.False(t, iter.HasNext())
}
