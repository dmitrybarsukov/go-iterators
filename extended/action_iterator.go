package extended

import "iterator/commons"

type actionIterator[T any] struct {
	innerIter commons.Iter[T]
	action    func(T)
}

func ActionIterator[T any](iter commons.Iter[T], action func(T)) commons.Iter[T] {
	if action == nil {
		panic(commons.ErrFuncIsNil)
	}
	return &actionIterator[T]{innerIter: iter, action: action}
}

func (i *actionIterator[T]) HasNext() bool {
	return i.innerIter.HasNext()
}

func (i *actionIterator[T]) Next() T {
	next := i.innerIter.Next()
	i.action(next)
	return next
}
