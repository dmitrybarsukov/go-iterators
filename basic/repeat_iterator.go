package basic

import "iterator/commons"

type repeatIterator[T any] struct {
	value T
	count int
}

func RepeatIterator[T any](value T, count int) commons.Iter[T] {
	if count < 0 {
		panic(commons.ErrArgumentCountNegative)
	}
	return &repeatIterator[T]{value: value, count: count}
}

func (i *repeatIterator[T]) HasNext() bool {
	return i.count > 0
}

func (i *repeatIterator[T]) Next() T {
	i.count -= 1
	return i.value
}
