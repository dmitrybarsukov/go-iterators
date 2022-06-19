package iterator

import (
	"iterator/basic"
	"iterator/commons"
	"iterator/extended"
)

type Iterator[T any] struct {
	iter commons.Iter[T]
}

func OfSlice[T any](slice []T) *Iterator[T] {
	return &Iterator[T]{iter: basic.Slice(slice)}
}

func OfMap[TK comparable, TV any](mapp map[TK]TV) *Iterator[basic.KeyValue[TK, TV]] {
	return &Iterator[basic.KeyValue[TK, TV]]{iter: basic.Map(mapp)}
}

func (i *Iterator[T]) HasNext() bool {
	return i.iter.HasNext()
}

func (i *Iterator[T]) Next() T {
	return i.iter.Next()
}

func (i *Iterator[T]) ToSlice() []T {
	result := make([]T, 0)
	for i.iter.HasNext() {
		result = append(result, i.iter.Next())
	}
	return result
}

func (i *Iterator[T]) Skip(count int) *Iterator[T] {
	return &Iterator[T]{iter: extended.SkippingIterator(i.iter, count)}
}

func (i *Iterator[T]) Limit(count int) *Iterator[T] {
	return &Iterator[T]{iter: extended.LimitingIterator(i.iter, count)}
}

//func (i *Iterator[TSrc]) Map[TRes any](mappingFunc extended.MappingFunc[TSrc, TRes]) *Iterator[TRes] {
//	return &Iterator[TRes]{
//		iter: extended.MappingIterator[TSrc, TRes](i.iter, mappingFunc),
//	}
//}
