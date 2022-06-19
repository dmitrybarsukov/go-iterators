package util

import "iterator/commons"

func ToSlice[T any](iter commons.Iter[T]) []T {
	result := make([]T, 0)
	for iter.HasNext() {
		result = append(result, iter.Next())
	}
	return result
}

func ToMap[TItem any, TKey comparable](
	iter commons.Iter[TItem],
	keyMappingFunc func(TItem) TKey,
) map[TKey]TItem {
	if keyMappingFunc == nil {
		panic(commons.ErrFuncIsNil)
	}
	result := make(map[TKey]TItem)
	for iter.HasNext() {
		item := iter.Next()
		result[keyMappingFunc(item)] = item
	}
	return result
}

func ToMapWithValue[TItem any, TKey comparable, TValue any](
	iter commons.Iter[TItem],
	keyMappingFunc func(TItem) TKey,
	valueMappingFunc func(TItem) TValue,
) map[TKey]TValue {
	if keyMappingFunc == nil || valueMappingFunc == nil {
		panic(commons.ErrFuncIsNil)
	}
	result := make(map[TKey]TValue)
	for iter.HasNext() {
		item := iter.Next()
		result[keyMappingFunc(item)] = valueMappingFunc(item)
	}
	return result
}

func Count[T any](iter commons.Iter[T]) int {
	result := 0
	for ; iter.HasNext(); result += 1 {
		_ = iter.Next()
	}
	return result
}

func CountMatching[T any](iter commons.Iter[T], predicate func(T) bool) int {
	if predicate == nil {
		panic(commons.ErrFuncIsNil)
	}
	result := 0
	for iter.HasNext() {
		if predicate(iter.Next()) {
			result += 1
		}
	}
	return result
}

func First[T any](iter commons.Iter[T]) (T, bool) {
	if !iter.HasNext() {
		return *new(T), false
	}
	return iter.Next(), true
}

func FirstOrZeroValue[T any](iter commons.Iter[T]) T {
	if item, ok := First(iter); ok {
		return item
	}
	return *new(T)
}

func FirstOrDefault[T any](iter commons.Iter[T], def T) T {
	if item, ok := First(iter); ok {
		return item
	}
	return def
}

func Last[T any](iter commons.Iter[T]) (T, bool) {
	var item T
	var exist = false
	for iter.HasNext() {
		exist = true
		item = iter.Next()
	}
	return item, exist
}

func LastOrZeroValue[T any](iter commons.Iter[T]) T {
	if item, ok := Last(iter); ok {
		return item
	}
	return *new(T)
}

func LastOrDefault[T any](iter commons.Iter[T], def T) T {
	if item, ok := Last(iter); ok {
		return item
	}
	return def
}
