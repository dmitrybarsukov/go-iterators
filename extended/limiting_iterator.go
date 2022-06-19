package extended

import "iterator/commons"

type limitingIterator[T any] struct {
	innerIter commons.Iter[T]
	count     int
}

func LimitingIterator[T any](iter commons.Iter[T], count int) commons.Iter[T] {
	return &limitingIterator[T]{innerIter: iter, count: count}
}

func (i *limitingIterator[T]) HasNext() bool {
	return i.count > 0 && i.innerIter.HasNext()
}

func (i *limitingIterator[T]) Next() T {
	i.count -= 1
	if i.count < 0 {
		panic(commons.ErrIterEnded)
	}
	return i.innerIter.Next()
}
