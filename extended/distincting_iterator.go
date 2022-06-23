package extended

import "iterator/commons"

type distinctingIterator[TItem any, TKey comparable] struct {
	innerIter  commons.Iter[TItem]
	keyFunc    func(TItem) TKey
	item       TItem
	key        TKey
	uniqueKeys map[TKey]struct{}
	ended      bool
}

func DistinctingIterator[TItem any, TKey comparable](iter commons.Iter[TItem], keyFunc func(TItem) TKey) commons.Iter[TItem] {
	if keyFunc == nil {
		panic(commons.ErrFuncIsNil)
	}
	return &distinctingIterator[TItem, TKey]{
		innerIter:  iter,
		keyFunc:    keyFunc,
		uniqueKeys: make(map[TKey]struct{}),
	}
}

func (i *distinctingIterator[TItem, TKey]) HasNext() bool {
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

func (i *distinctingIterator[TItem, TKey]) Next() TItem {
	if i.ended {
		panic(commons.ErrIterEnded)
	}
	i.uniqueKeys[i.key] = struct{}{}
	return i.item
}
