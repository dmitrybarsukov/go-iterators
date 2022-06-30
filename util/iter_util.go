package util

import "iterator/commons"

func Self[T any](i T) T {
	return i
}

func SelfAny[T any](i T) any {
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

func ToMapMultiple[TItem any, TKey comparable](
	iter commons.Iter[TItem],
	keyMappingFunc func(TItem) TKey,
) map[TKey][]TItem {
	if keyMappingFunc == nil {
		panic(commons.ErrFuncIsNil)
	}
	result := make(map[TKey][]TItem)
	for iter.HasNext() {
		item := iter.Next()
		result[keyMappingFunc(item)] = append(result[keyMappingFunc(item)], item)
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

func ToReceiveChannel[T any](iter commons.Iter[T]) <-chan T {
	channel := make(chan T)
	go WriteToChannelAndClose(iter, channel)
	return channel
}

func WriteToChannel[T any](iter commons.Iter[T], channel chan<- T) {
	ForEach(iter, func(item T) { channel <- item })
}

func WriteToChannelAndClose[T any](iter commons.Iter[T], channel chan<- T) {
	WriteToChannel(iter, channel)
	close(channel)
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
	if keyFunc == nil {
		panic(commons.ErrFuncIsNil)
	}
	if !iter.HasNext() {
		return
	}
	ok = true
	result = iter.Next()
	key := keyFunc(result)
	for iter.HasNext() {
		value := iter.Next()
		newKey := keyFunc(value)
		if newKey > key {
			result = value
			key = newKey
		}
	}
	return
}

func MaxByOrZeroValue[TItem any, TKey commons.Ordered](iter commons.Iter[TItem], keyFunc func(TItem) TKey) (result TItem) {
	if keyFunc == nil {
		panic(commons.ErrFuncIsNil)
	}
	if item, ok := MaxBy(iter, keyFunc); ok {
		result = item
	}
	return
}

func MaxByOrDefault[TItem any, TKey commons.Ordered](iter commons.Iter[TItem], keyFunc func(TItem) TKey, def TItem) TItem {
	if keyFunc == nil {
		panic(commons.ErrFuncIsNil)
	}
	if item, ok := MaxBy(iter, keyFunc); ok {
		return item
	}
	return def
}

func MinBy[TItem any, TKey commons.Ordered](iter commons.Iter[TItem], keyFunc func(TItem) TKey) (result TItem, ok bool) {
	if keyFunc == nil {
		panic(commons.ErrFuncIsNil)
	}
	if !iter.HasNext() {
		return
	}
	ok = true
	result = iter.Next()
	key := keyFunc(result)
	for iter.HasNext() {
		value := iter.Next()
		newKey := keyFunc(value)
		if newKey < key {
			result = value
			key = newKey
		}
	}
	return
}

func MinByOrZeroValue[TItem any, TKey commons.Ordered](iter commons.Iter[TItem], keyFunc func(TItem) TKey) (result TItem) {
	if keyFunc == nil {
		panic(commons.ErrFuncIsNil)
	}
	if item, ok := MinBy(iter, keyFunc); ok {
		result = item
	}
	return
}

func MinByOrDefault[TItem any, TKey commons.Ordered](iter commons.Iter[TItem], keyFunc func(TItem) TKey, def TItem) TItem {
	if keyFunc == nil {
		panic(commons.ErrFuncIsNil)
	}
	if item, ok := MinBy(iter, keyFunc); ok {
		return item
	}
	return def
}

func ForEach[T any](iter commons.Iter[T], action func(T)) {
	if action == nil {
		panic(commons.ErrFuncIsNil)
	}
	for iter.HasNext() {
		action(iter.Next())
	}
}

func ForEachCollectingErrors[T any](iter commons.Iter[T], action func(T) error) (result []error) {
	if action == nil {
		panic(commons.ErrFuncIsNil)
	}
	for iter.HasNext() {
		err := action(iter.Next())
		if err != nil {
			result = append(result, err)
		}
	}
	return
}

func ForEachUntilFirstError[T any](iter commons.Iter[T], action func(T) error) error {
	if action == nil {
		panic(commons.ErrFuncIsNil)
	}
	for iter.HasNext() {
		err := action(iter.Next())
		if err != nil {
			return err
		}
	}
	return nil
}

func All[T any](iter commons.Iter[T], predicate func(T) bool) bool {
	if predicate == nil {
		panic(commons.ErrFuncIsNil)
	}
	for iter.HasNext() {
		if !predicate(iter.Next()) {
			return false
		}
	}
	return true
}

func Any[T any](iter commons.Iter[T], predicate func(T) bool) bool {
	if predicate == nil {
		panic(commons.ErrFuncIsNil)
	}
	for iter.HasNext() {
		if predicate(iter.Next()) {
			return true
		}
	}
	return false
}

func None[T any](iter commons.Iter[T], predicate func(T) bool) bool {
	if predicate == nil {
		panic(commons.ErrFuncIsNil)
	}
	for iter.HasNext() {
		if predicate(iter.Next()) {
			return false
		}
	}
	return true
}

func SumBy[TItem any, TValue commons.Ordered](iter commons.Iter[TItem], valueFunc func(TItem) TValue) (sum TValue) {
	ForEach(iter, func(item TItem) {
		sum += valueFunc(item)
	})
	return sum
}
