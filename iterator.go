package iterator

import (
	"iterator/basic"
	"iterator/commons"
	"iterator/extended"
	"iterator/util"
)

type Iterator[T any] struct {
	iter commons.Iter[T]
}

func Slice[T any](slice []T) Iterator[T] {
	return Iterator[T]{iter: basic.SliceIterator(slice)}
}

func Map[TK comparable, TV any](mapp map[TK]TV) Iterator[commons.KeyValue[TK, TV]] {
	return Iterator[commons.KeyValue[TK, TV]]{iter: basic.MapIterator(mapp)}
}

func IntRange(startInclusive, endExclusive, step int) commons.Iter[int] {
	return Iterator[int]{iter: basic.IntRangeIterator(startInclusive, endExclusive, step)}
}

func Generator[T any](generateFunc func(int) T) commons.Iter[T] {
	return Iterator[T]{iter: basic.GeneratorIterator(generateFunc)}
}

func Repeat[T any](value T, count int) commons.Iter[T] {
	return Iterator[T]{iter: basic.RepeatIterator(value, count)}
}

func (i Iterator[T]) HasNext() bool {
	return i.iter.HasNext()
}

func (i Iterator[T]) Next() T {
	return i.iter.Next()
}

func (i Iterator[T]) ToSlice() []T {
	return util.ToSlice(i.iter)
}

func (i Iterator[T]) Count() int {
	return util.Count(i.iter)
}

func (i Iterator[T]) CountMatching(predicate func(T) bool) int {
	return util.CountMatching(i.iter, predicate)
}

func (i Iterator[T]) First() (T, bool) {
	return util.First(i.iter)
}

func (i Iterator[T]) FirstOrZeroValue() T {
	return util.FirstOrZeroValue(i.iter)
}

func (i Iterator[T]) FirstOrDefault(def T) T {
	return util.FirstOrDefault(i.iter, def)
}

func (i Iterator[T]) Last() (T, bool) {
	return util.Last(i.iter)
}

func (i Iterator[T]) LastOrZeroValue() T {
	return util.LastOrZeroValue(i.iter)
}

func (i Iterator[T]) LastOrDefault(def T) T {
	return util.LastOrDefault(i.iter, def)
}

func (i Iterator[T]) Skip(count int) Iterator[T] {
	return Iterator[T]{iter: extended.SkippingIterator(i.iter, count)}
}

func (i Iterator[T]) Limit(count int) Iterator[T] {
	return Iterator[T]{iter: extended.LimitingIterator(i.iter, count)}
}

func (i Iterator[T]) Filter(predicate func(T) bool) Iterator[T] {
	return Iterator[T]{iter: extended.FilteringIterator(i.iter, predicate)}
}

func (i Iterator[T]) SortedByInt(keyFunc func(T) int) Iterator[T] {
	return Iterator[T]{iter: extended.SortingIterator(i.iter, keyFunc, false)}
}

func (i Iterator[T]) SortedByIntDescending(keyFunc func(T) int) Iterator[T] {
	return Iterator[T]{iter: extended.SortingIterator(i.iter, keyFunc, true)}
}

func (i Iterator[T]) SortedByInt64(keyFunc func(T) int64) Iterator[T] {
	return Iterator[T]{iter: extended.SortingIterator(i.iter, keyFunc, false)}
}

func (i Iterator[T]) SortedByInt64Descending(keyFunc func(T) int64) Iterator[T] {
	return Iterator[T]{iter: extended.SortingIterator(i.iter, keyFunc, true)}
}

func (i Iterator[T]) SortedByString(keyFunc func(T) string) Iterator[T] {
	return Iterator[T]{iter: extended.SortingIterator(i.iter, keyFunc, false)}
}

func (i Iterator[T]) SortedByStringDescending(keyFunc func(T) string) Iterator[T] {
	return Iterator[T]{iter: extended.SortingIterator(i.iter, keyFunc, true)}
}

func (i Iterator[T]) SortedByFloat32(keyFunc func(T) float32) Iterator[T] {
	return Iterator[T]{iter: extended.SortingIterator(i.iter, keyFunc, false)}
}

func (i Iterator[T]) SortedByFloat32Descending(keyFunc func(T) float32) Iterator[T] {
	return Iterator[T]{iter: extended.SortingIterator(i.iter, keyFunc, true)}
}

func (i Iterator[T]) SortedByFloat64(keyFunc func(T) float64) Iterator[T] {
	return Iterator[T]{iter: extended.SortingIterator(i.iter, keyFunc, false)}
}

func (i Iterator[T]) SortedByFloat64Descending(keyFunc func(T) float64) Iterator[T] {
	return Iterator[T]{iter: extended.SortingIterator(i.iter, keyFunc, true)}
}

//func (i Iterator[TSrc]) MapIterator[TRes any](mappingFunc extended.MappingFunc[TSrc, TRes]) Iterator[TRes] {
//	return &Iterator[TRes]{
//		iter: extended.MappingIterator[TSrc, TRes](i.iter, mappingFunc),
//	}
//}
