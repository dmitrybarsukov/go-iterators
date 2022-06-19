package basic

import (
	"testing"

	"iterator/commons"
	"iterator/util"

	"github.com/stretchr/testify/assert"
)

func TestMapIteratorEmpty(t *testing.T) {
	items := make(map[int]int)
	iter := MapIterator(items)
	count := 0
	for iter.HasNext() {
		_ = iter.Next()
		count += 1
	}
	assert.Equal(t, 0, count)
}

func TestMapIteratorMany(t *testing.T) {
	items := map[int]string{1: "1", 2: "2", 3: "3"}
	iter := MapIterator(items)
	result := util.ToMapKeyValue(iter)
	assert.Equal(t, items, result)
}

func TestMapIteratorPanicsIfIteratorEnded(t *testing.T) {
	items := map[int]string{1: "1", 2: "2", 3: "3"}
	iter := MapIterator(items)
	for iter.HasNext() {
		_ = iter.Next()
	}
	assert.PanicsWithValue(t, commons.ErrIterEnded, func() {
		iter.Next()
	})
}

func TestMapIteratorNil(t *testing.T) {
	var items map[int]int
	iter := MapIterator(items)
	assert.False(t, iter.HasNext())
}
