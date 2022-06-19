package basic

import "iterator/commons"

type sliceIterator[T any] struct {
	slice        []T
	currentIndex int
}

func SliceIterator[T any](slice []T) commons.Iter[T] {
	return &sliceIterator[T]{slice: slice}
}

func (i *sliceIterator[T]) HasNext() bool {
	return i.currentIndex < len(i.slice)
}

func (i *sliceIterator[T]) Next() T {
	if i.currentIndex >= len(i.slice) {
		panic(commons.ErrIterEnded)
	}
	i.currentIndex += 1
	return i.slice[i.currentIndex-1]
}
