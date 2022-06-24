package basic

import "iterator/commons"

type chanIterator[T any] struct {
	channel  <-chan T
	nextItem T
	ended    bool
}

func ChanIterator[T any](channel <-chan T) commons.Iter[T] {
	return &chanIterator[T]{channel: channel}
}

func (i *chanIterator[T]) HasNext() bool {
	nextItem, hasNext := <-i.channel
	if !hasNext {
		i.ended = true
	}
	i.nextItem = nextItem
	return hasNext
}

func (i *chanIterator[T]) Next() T {
	if i.ended {
		panic(commons.ErrIterEnded)
	}
	return i.nextItem
}
