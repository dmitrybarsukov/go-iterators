package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratorIterator(t *testing.T) {
	iter := GeneratorIterator(func(i int) int { return i * 2 })
	result := make([]int, 0)
	for i := 0; i < 5; i++ {
		assert.True(t, iter.HasNext())
		result = append(result, iter.Next())
	}
	assert.Equal(t, []int{0, 2, 4, 6, 8}, result)
}
