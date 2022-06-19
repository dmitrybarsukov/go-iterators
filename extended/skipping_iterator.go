package extended

import "iterator/commons"

type skippingIterator[T any] struct {
	innerIter commons.Iter[T]
	count     int
}

func SkippingIterator[T any](iter commons.Iter[T], count int) commons.Iter[T] {
	return &skippingIterator[T]{innerIter: iter, count: count}
}

func (i *skippingIterator[T]) checkSkipped() {
	for i.count > 0 && i.innerIter.HasNext() {
		i.count -= 1
		_ = i.innerIter.Next()
	}
}

func (i *skippingIterator[T]) HasNext() bool {
	i.checkSkipped()
	return i.innerIter.HasNext()
}

func (i *skippingIterator[T]) Next() T {
	i.checkSkipped()
	return i.innerIter.Next()
}
