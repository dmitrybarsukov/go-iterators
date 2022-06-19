package extended

import (
	"iterator/commons"
)

type Predicate[T any] func(T) bool

type filteringIterator[T any] struct {
	innerIter commons.Iter[T]
	predicate Predicate[T]
	item      T
	ended     bool
}

func FilteringIterator[T any](iter commons.Iter[T], predicate Predicate[T]) commons.Iter[T] {
	if predicate == nil {
		panic(commons.ErrFuncIsNil)
	}
	return &filteringIterator[T]{innerIter: iter, predicate: predicate}
}

func (i *filteringIterator[T]) HasNext() bool {
	for i.innerIter.HasNext() {
		i.item = i.innerIter.Next()
		if i.predicate(i.item) {
			return true
		}
	}
	i.ended = true
	return false
}

func (i filteringIterator[T]) Next() T {
	if i.ended {
		panic(commons.ErrIterEnded)
	}
	return i.item
}
