package basic

import "iterator/commons"

type mapIterator[TK comparable, TV any] struct {
	mapp         map[TK]TV
	keys         []TK
	currentIndex int
}

func MapIterator[TK comparable, TV any](mapp map[TK]TV) commons.Iter[commons.KeyValue[TK, TV]] {
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

func (i *mapIterator[TK, TV]) Next() commons.KeyValue[TK, TV] {
	i.checkKeysInitialized()
	if i.currentIndex >= len(i.keys) {
		panic(commons.ErrIterEnded)
	}
	i.currentIndex += 1
	key := i.keys[i.currentIndex-1]
	value := i.mapp[key]
	return commons.KeyValue[TK, TV]{key, value}
}
