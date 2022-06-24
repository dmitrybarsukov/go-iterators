package extended

import "iterator/commons"

type distinctingIterator[T any] struct {
	innerIter  commons.Iter[T]
	keyFunc    func(T) any
	item       T
	key        any
	uniqueKeys map[any]struct{}
	ended      bool
}

func DistinctingIterator[T any](iter commons.Iter[T], keyFunc func(T) any) commons.Iter[T] {
	if keyFunc == nil {
		panic(commons.ErrFuncIsNil)
	}
	return &distinctingIterator[T]{
		innerIter:  iter,
		keyFunc:    keyFunc,
		uniqueKeys: make(map[any]struct{}),
	}
}

func (i *distinctingIterator[T]) HasNext() bool {
	for i.innerIter.HasNext() {
		i.item = i.innerIter.Next()
		i.key = i.keyFunc(i.item)
		if _, ok := i.uniqueKeys[i.key]; !ok {
			return true
		}
	}
	i.ended = true
	return false
}

func (i *distinctingIterator[T]) Next() T {
	if i.ended {
		panic(commons.ErrIterEnded)
	}
	i.uniqueKeys[i.key] = struct{}{}
	return i.item
}
