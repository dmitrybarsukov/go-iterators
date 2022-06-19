package basic

import (
	"go-iterators"
	"go-iterators/errors"
)

type sliceIterator[T any] struct {
	slice        []T
	currentIndex int
}

func Slice[T any](slice []T) go_iterators.Iter[T] {
	return &sliceIterator[T]{slice: slice}
}

func (i *sliceIterator[T]) HasNext() bool {
	return i.currentIndex < len(i.slice)
}

func (i *sliceIterator[T]) Next() T {
	if i.currentIndex >= len(i.slice) {
		panic(errors.ErrIterEnded)
	}
	i.currentIndex += 1
	return i.slice[i.currentIndex-1]
}
