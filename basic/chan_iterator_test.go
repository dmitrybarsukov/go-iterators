package basic

import (
	"testing"

	"iterator/util"

	"github.com/stretchr/testify/assert"
)

func TestChanIterator(t *testing.T) {
	channel := make(chan int)
	iter := ChanIterator(channel)
	go func() {
		channel <- 1
		channel <- 2
		channel <- 3
		close(channel)
	}()
	items := util.ToSlice(iter)
	assert.Equal(t, []int{1, 2, 3}, items)
}
