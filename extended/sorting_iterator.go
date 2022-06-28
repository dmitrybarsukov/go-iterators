package extended

import (
	"sort"

	"iterator/basic"
	"iterator/commons"
	"iterator/util"
	"iterator/util/comparator"
)

type sortingIterator[T any] struct {
	innerIter     commons.Iter[T]
	compareFunc   func(T, T) bool
	keyFunc       func(T) any
	isDesc        bool
	isInitialized bool
}

func SortingIterator[T any](iter commons.Iter[T], compareFunc func(T, T) bool) commons.Iter[T] {
	if compareFunc == nil {
		panic(commons.ErrFuncIsNil)
	}
	return &sortingIterator[T]{innerIter: iter, compareFunc: compareFunc}
}

func SortingIteratorAsc[T any](iter commons.Iter[T], keyFunc func(T) any) commons.Iter[T] {
	if keyFunc == nil {
		panic(commons.ErrFuncIsNil)
	}
	return &sortingIterator[T]{innerIter: iter, keyFunc: keyFunc, isDesc: false}
}

func SortingIteratorDesc[T any](iter commons.Iter[T], keyFunc func(T) any) commons.Iter[T] {
	if keyFunc == nil {
		panic(commons.ErrFuncIsNil)
	}
	return &sortingIterator[T]{innerIter: iter, keyFunc: keyFunc, isDesc: true}
}

func (i *sortingIterator[T]) checkInitialized() {
	if !i.isInitialized {
		i.isInitialized = true
		items := util.ToSlice(i.innerIter)

		var compFunc func(int, int) bool
		if i.compareFunc != nil {
			compFunc = func(j, k int) bool { return i.compareFunc(items[j], items[k]) }
		} else {
			uniComp := comparator.NewUniversalComparator()
			if i.isDesc {
				compFunc = func(j, k int) bool { return !uniComp.Compare(i.keyFunc(items[j]), i.keyFunc(items[k])) }
			} else {
				compFunc = func(j, k int) bool { return uniComp.Compare(i.keyFunc(items[j]), i.keyFunc(items[k])) }
			}
		}

		sort.Slice(items, compFunc)
		i.innerIter = basic.SliceIterator(items)
	}
}

func (i *sortingIterator[T]) HasNext() bool {
	i.checkInitialized()
	return i.innerIter.HasNext()
}

func (i *sortingIterator[T]) Next() T {
	i.checkInitialized()
	return i.innerIter.Next()
}
