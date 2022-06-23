package extended

import (
	"sort"

	"iterator/basic"
	"iterator/commons"
	"iterator/util"
)

type sortingIterator[TItem any] struct {
	innerIter     commons.Iter[TItem]
	compareFunc   func(TItem, TItem) bool
	isInitialized bool
}

func SortingIterator[TItem any](iter commons.Iter[TItem], compareFunc func(TItem, TItem) bool) commons.Iter[TItem] {
	if compareFunc == nil {
		panic(commons.ErrFuncIsNil)
	}
	return &sortingIterator[TItem]{innerIter: iter, compareFunc: compareFunc}
}

func SortingIteratorAsc[TItem any, TKey commons.Ordered](iter commons.Iter[TItem], keyFunc func(TItem) TKey) commons.Iter[TItem] {
	if keyFunc == nil {
		panic(commons.ErrFuncIsNil)
	}
	return SortingIterator[TItem](iter, func(item1, item2 TItem) bool { return keyFunc(item1) < keyFunc(item2) })
}

func SortingIteratorDesc[TItem any, TKey commons.Ordered](iter commons.Iter[TItem], keyFunc func(TItem) TKey) commons.Iter[TItem] {
	if keyFunc == nil {
		panic(commons.ErrFuncIsNil)
	}
	return SortingIterator[TItem](iter, func(item1, item2 TItem) bool { return keyFunc(item1) > keyFunc(item2) })
}

func (i *sortingIterator[TItem]) checkInitialized() {
	if !i.isInitialized {
		i.isInitialized = true
		items := util.ToSlice(i.innerIter)

		sort.Slice(items, func(j, k int) bool { return i.compareFunc(items[j], items[k]) })
		i.innerIter = basic.SliceIterator(items)
	}
}

func (i *sortingIterator[TItem]) HasNext() bool {
	i.checkInitialized()
	return i.innerIter.HasNext()
}

func (i *sortingIterator[TItem]) Next() TItem {
	i.checkInitialized()
	return i.innerIter.Next()
}
