package basic

import (
	"testing"

	"iterator/commons"

	"github.com/stretchr/testify/assert"
)

func TestMapIteratorEmpty(t *testing.T) {
	items := make(map[int]int)
	iter := Map(items)
	count := 0
	for iter.HasNext() {
		_ = iter.Next()
		count += 1
	}
	assert.Equal(t, 0, count)
}

func TestMapIteratorMany(t *testing.T) {
	items := map[int]string{1: "1", 2: "2", 3: "3"}
	iter := Map(items)
	result := make(map[int]string)
	for iter.HasNext() {
		key, value := iter.Next().Pair()
		result[key] = value
	}
	assert.Equal(t, items, result)
}

func TestMapIteratorPanicsIfIteratorEnded(t *testing.T) {
	items := map[int]string{1: "1", 2: "2", 3: "3"}
	iter := Map(items)
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

func TestMapIteratorNil(t *testing.T) {
	var items map[int]int
	iter := Map(items)
	assert.False(t, iter.HasNext())
}
