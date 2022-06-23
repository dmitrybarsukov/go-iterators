package extended

import (
	"iterator/basic"
	"iterator/commons"
)

type prependingAppendingIterator[T any] struct {
	innerIter      commons.Iter[T]
	itemsToAppend  []T
	itemsToPrepend []T
	iters          []commons.Iter[T]
	index          int
}

func AppendingIterator[T any](iter commons.Iter[T], items ...T) commons.Iter[T] {
	if appIter, ok := iter.(*prependingAppendingIterator[T]); ok {
		appIter.itemsToAppend = append(appIter.itemsToAppend, items...)
		return appIter
	}
	return &prependingAppendingIterator[T]{innerIter: iter, itemsToAppend: items}
}

func PrependingIterator[T any](iter commons.Iter[T], items ...T) commons.Iter[T] {
	if appIter, ok := iter.(*prependingAppendingIterator[T]); ok {
		appIter.itemsToPrepend = append(items, appIter.itemsToPrepend...)
		return appIter
	}
	return &prependingAppendingIterator[T]{innerIter: iter, itemsToPrepend: items}
}

func (i *prependingAppendingIterator[T]) checkInitialized() {
	if i.iters == nil {
		i.iters = []commons.Iter[T]{
			basic.SliceIterator(i.itemsToPrepend),
			i.innerIter,
			basic.SliceIterator(i.itemsToAppend),
		}
	}
}

func (i *prependingAppendingIterator[T]) HasNext() bool {
	i.checkInitialized()
	hasNext := i.iters[i.index].HasNext()
	if !hasNext {
		if i.index == 2 {
			return false
		}
		i.index += 1
		hasNext = i.iters[i.index].HasNext()
	}
	return hasNext
}

func (i *prependingAppendingIterator[T]) Next() T {
	i.checkInitialized()
	return i.iters[i.index].Next()
}
