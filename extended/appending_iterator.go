package extended

import (
	"iterator/basic"
	"iterator/commons"
)

type appendingIterator[T any] struct {
	innerIter            commons.Iter[T]
	itemsToAppend        []T
	primaryIterExhausted bool
}

func AppendingIterator[T any](iter commons.Iter[T], items ...T) commons.Iter[T] {
	if appIter, ok := iter.(*appendingIterator[T]); ok {
		appIter.itemsToAppend = append(appIter.itemsToAppend, items...)
		return appIter
	}
	return &appendingIterator[T]{innerIter: iter, itemsToAppend: items}
}

func (i *appendingIterator[T]) HasNext() bool {
	hasNext := i.innerIter.HasNext()
	if !i.primaryIterExhausted && !hasNext {
		i.primaryIterExhausted = true
		i.innerIter = basic.SliceIterator(i.itemsToAppend)
		hasNext = i.innerIter.HasNext()
	}
	return hasNext
}

func (i *appendingIterator[T]) Next() T {
	return i.innerIter.Next()
}
