package extended

import "iterator/commons"

type mappingIterator[TSrc, TRes any] struct {
	innerIter   commons.Iter[TSrc]
	mappingFunc func(TSrc) TRes
}

func MappingIterator[TSrc, TRes any](iter commons.Iter[TSrc], mappingFunc func(TSrc) TRes) commons.Iter[TRes] {
	if mappingFunc == nil {
		panic(commons.ErrFuncIsNil)
	}
	return &mappingIterator[TSrc, TRes]{
		innerIter:   iter,
		mappingFunc: mappingFunc,
	}
}

func (m *mappingIterator[TSrc, TRes]) HasNext() bool {
	return m.innerIter.HasNext()
}

func (m *mappingIterator[TSrc, TRes]) Next() TRes {
	return m.mappingFunc(m.innerIter.Next())
}
