package extended

import (
	"sort"

	"iterator/basic"
	"iterator/commons"
	"iterator/util"
)

type sortingIterator[TItem any, TKey commons.Ordered] struct {
	innerIter     commons.Iter[TItem]
	keyFunc       func(TItem) TKey
	isDescending  bool
	isInitialized bool
}

func SortingIterator[TItem any, TKey commons.Ordered](iter commons.Iter[TItem], keyFunc func(TItem) TKey, isDescending bool) commons.Iter[TItem] {
	return &sortingIterator[TItem, TKey]{innerIter: iter, keyFunc: keyFunc, isDescending: isDescending}
}

func (i *sortingIterator[TItem, TKey]) checkInitialized() {
	if !i.isInitialized {
		i.isInitialized = true
		items := util.ToSlice(i.innerIter)
		var sortFunc func(j, k int) bool
		if i.isDescending {
			sortFunc = func(j, k int) bool {
				return i.keyFunc(items[j]) > i.keyFunc(items[k])
			}
		} else {
			sortFunc = func(j, k int) bool {
				return i.keyFunc(items[j]) < i.keyFunc(items[k])
			}
		}
		sort.Slice(items, sortFunc)
		i.innerIter = basic.SliceIterator(items)
	}
}

func (i *sortingIterator[TItem, TKey]) HasNext() bool {
	i.checkInitialized()
	return i.innerIter.HasNext()
}

func (i *sortingIterator[TItem, TKey]) Next() TItem {
	i.checkInitialized()
	return i.innerIter.Next()
}
