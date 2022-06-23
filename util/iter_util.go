package util

import "iterator/commons"

func Self[T any](i T) T {
	return i
}

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

func ToMapKeyValue[TKey comparable, TValue any](iter commons.Iter[commons.KeyValue[TKey, TValue]]) map[TKey]TValue {
	result := make(map[TKey]TValue)
	for iter.HasNext() {
		key, value := iter.Next().Pair()
		result[key] = value
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

func First[T any](iter commons.Iter[T]) (result T, ok bool) {
	if iter.HasNext() {
		result = iter.Next()
		ok = true
	}
	return
}

func FirstOrZeroValue[T any](iter commons.Iter[T]) (result T) {
	if item, ok := First(iter); ok {
		result = item
	}
	return
}

func FirstOrDefault[T any](iter commons.Iter[T], def T) T {
	if item, ok := First(iter); ok {
		return item
	}
	return def
}

func Last[T any](iter commons.Iter[T]) (result T, ok bool) {
	for iter.HasNext() {
		ok = true
		result = iter.Next()
	}
	return
}

func LastOrZeroValue[T any](iter commons.Iter[T]) (result T) {
	if item, ok := Last(iter); ok {
		result = item
	}
	return
}

func LastOrDefault[T any](iter commons.Iter[T], def T) T {
	if item, ok := Last(iter); ok {
		return item
	}
	return def
}

func MaxBy[TItem any, TKey commons.Ordered](iter commons.Iter[TItem], keyFunc func(TItem) TKey) (result TItem, ok bool) {
	if !iter.HasNext() {
		return
	}
	ok = true
	result = iter.Next()
	for iter.HasNext() {
		value := iter.Next()
		if keyFunc(value) > keyFunc(result) {
			result = value
		}
	}
	return
}

func MaxByOrZeroValue[TItem any, TKey commons.Ordered](iter commons.Iter[TItem], keyFunc func(TItem) TKey) (result TItem) {
	if item, ok := MaxBy(iter, keyFunc); ok {
		result = item
	}
	return
}

func MaxByOrDefault[TItem any, TKey commons.Ordered](iter commons.Iter[TItem], keyFunc func(TItem) TKey, def TItem) TItem {
	if item, ok := MaxBy(iter, keyFunc); ok {
		return item
	}
	return def
}

func MinBy[TItem any, TKey commons.Ordered](iter commons.Iter[TItem], keyFunc func(TItem) TKey) (result TItem, ok bool) {
	if !iter.HasNext() {
		return
	}
	ok = true
	result = iter.Next()
	for iter.HasNext() {
		value := iter.Next()
		if keyFunc(value) < keyFunc(result) {
			result = value
		}
	}
	return
}

func MinByOrZeroValue[TItem any, TKey commons.Ordered](iter commons.Iter[TItem], keyFunc func(TItem) TKey) (result TItem) {
	if item, ok := MinBy(iter, keyFunc); ok {
		result = item
	}
	return
}

func MinByOrDefault[TItem any, TKey commons.Ordered](iter commons.Iter[TItem], keyFunc func(TItem) TKey, def TItem) TItem {
	if item, ok := MinBy(iter, keyFunc); ok {
		return item
	}
	return def
}

func ForEach[TItem any](iter commons.Iter[TItem], action func(TItem)) {
	for iter.HasNext() {
		action(iter.Next())
	}
}

func ForEachCollectingErrors[TItem any](iter commons.Iter[TItem], action func(TItem) error) (result []error) {
	for iter.HasNext() {
		err := action(iter.Next())
		if err != nil {
			result = append(result, err)
		}
	}
	return
}

func ForEachUntilFirstError[TItem any](iter commons.Iter[TItem], action func(TItem) error) error {
	for iter.HasNext() {
		err := action(iter.Next())
		if err != nil {
			return err
		}
	}
	return nil
}
