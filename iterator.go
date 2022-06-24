package iterator

//go:generate go run generate/extensions.go

import (
	"iterator/basic"
	"iterator/commons"
	"iterator/extended"
	"iterator/util"
)

type Iterator[T any] struct {
	iter commons.Iter[T]
}

func Empty[T any]() Iterator[T] {
	return Iterator[T]{iter: basic.SliceIterator[T](nil)}
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

func (i Iterator[T]) Reversed() Iterator[T] {
	return Iterator[T]{iter: extended.ReversingIterator(i.iter)}
}

func (i Iterator[T]) Append(items ...T) Iterator[T] {
	return Iterator[T]{iter: extended.AppendingIterator(i.iter, items...)}
}

func (i Iterator[T]) Prepend(items ...T) Iterator[T] {
	return Iterator[T]{iter: extended.PrependingIterator(i.iter, items...)}
}

func (i Iterator[T]) OnEach(iter commons.Iter[T], action func(T)) commons.Iter[T] {
	return Iterator[T]{iter: extended.ActionIterator(iter, action)}
}

func (i Iterator[T]) ForEach(action func(T)) {
	util.ForEach(i.iter, action)
}

func (i Iterator[T]) ForEachCollectingErrors(action func(T) error) []error {
	return util.ForEachCollectingErrors(i.iter, action)
}

func (i Iterator[T]) ForEachUntilFirstError(action func(T) error) error {
	return util.ForEachUntilFirstError(i.iter, action)
}

func (i Iterator[T]) All(predicate func(T) bool) bool {
	return util.All(i.iter, predicate)
}

func (i Iterator[T]) Any(predicate func(T) bool) bool {
	return util.Any(i.iter, predicate)
}

func (i Iterator[T]) None(predicate func(T) bool) bool {
	return util.None(i.iter, predicate)
}

//func (i Iterator[TSrc]) MapIterator[TRes any](mappingFunc extended.MappingFunc[TSrc, TRes]) Iterator[TRes] {
//	return &Iterator[TRes]{
//		iter: extended.MappingIterator[TSrc, TRes](i.iter, mappingFunc),
//	}
//}
