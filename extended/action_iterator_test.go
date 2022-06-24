package extended

import (
	"testing"

	"iterator/basic"

	"github.com/stretchr/testify/assert"
)

func TestActionIterator(t *testing.T) {
	var values []int
	iter := ActionIterator(basic.SliceIterator([]int{1, 2, 3}), func(i int) {
		values = append(values, i)
	})
	assert.Empty(t, values)
	for iter.HasNext() {
		iter.Next()
	}
	assert.Equal(t, []int{1, 2, 3}, values)
}
