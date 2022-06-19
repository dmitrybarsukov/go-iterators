package extended

import (
	"iterator/basic"
	"iterator/commons"
	"iterator/util"
)

type reversingIterator[T any] struct {
	innerIter     commons.Iter[T]
	isInitialized bool
}

func ReversingIterator[T any](iter commons.Iter[T]) commons.Iter[T] {
	return &reversingIterator[T]{innerIter: iter}
}

func (i *reversingIterator[T]) checkInitialized() {
	if !i.isInitialized {
		i.isInitialized = true
		items := util.ToSlice(i.innerIter)
		count := len(items)
		result := make([]T, count)
		for j, item := range items {
			result[count-j-1] = item
		}
		i.innerIter = basic.SliceIterator(result)
	}
}

func (i *reversingIterator[T]) HasNext() bool {
	i.checkInitialized()
	return i.innerIter.HasNext()
}

func (i *reversingIterator[T]) Next() T {
	i.checkInitialized()
	return i.innerIter.Next()
}
