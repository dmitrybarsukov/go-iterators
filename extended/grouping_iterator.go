package extended

import "iterator/commons"

type groupingIterator[T any] struct {
	innerIter commons.Iter[T]
	keyFunc   func(T) any
	lastItem  T
	started   bool
	ended     bool
}

func GroupingIterator[T any](iter commons.Iter[T], keyFunc func(T) any) commons.Iter[[]T] {
	if keyFunc == nil {
		panic(commons.ErrFuncIsNil)
	}
	return &groupingIterator[T]{innerIter: SortingIteratorAsc(iter, keyFunc), keyFunc: keyFunc}
}

func (i *groupingIterator[T]) HasNext() bool {
	return i.innerIter.HasNext() || !i.ended
}

func (i *groupingIterator[T]) Next() []T {
	if !i.started {
		i.started = true
		i.lastItem = i.innerIter.Next()
	}
	items := []T{i.lastItem}
	currentKey := i.keyFunc(i.lastItem)
	for i.innerIter.HasNext() {
		i.lastItem = i.innerIter.Next()
		if i.keyFunc(i.lastItem) == currentKey {
			items = append(items, i.lastItem)
		} else {
			return items
		}
	}
	i.ended = true
	return items
}
