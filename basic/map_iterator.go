package basic

import (
	"go-iterators"
	"go-iterators/errors"
)

type keyValue[TK comparable, TV any] struct {
	Key   TK
	Value TV
}

func (v keyValue[TK, TV]) Pair() (TK, TV) {
	return v.Key, v.Value
}

type mapIterator[TK comparable, TV any] struct {
	mapp         map[TK]TV
	keys         []TK
	currentIndex int
}

func Map[TK comparable, TV any](mapp map[TK]TV) go_iterators.Iter[keyValue[TK, TV]] {
	return &mapIterator[TK, TV]{mapp: mapp}
}

func (i *mapIterator[TK, TV]) checkKeysInitialized() {
	if i.keys == nil {
		i.keys = make([]TK, 0, len(i.mapp))
		for k, _ := range i.mapp {
			i.keys = append(i.keys, k)
		}
	}
}

func (i *mapIterator[TK, TV]) HasNext() bool {
	i.checkKeysInitialized()
	return i.currentIndex < len(i.keys)
}

func (i *mapIterator[TK, TV]) Next() keyValue[TK, TV] {
	i.checkKeysInitialized()
	if i.currentIndex >= len(i.keys) {
		panic(errors.ErrIterEnded)
	}
	i.currentIndex += 1
	key := i.keys[i.currentIndex-1]
	value := i.mapp[key]
	return keyValue[TK, TV]{key, value}
}
