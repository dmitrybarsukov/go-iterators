package basic

import "iterator/commons"

type generatorIterator[T any] struct {
	index        int
	generateFunc func(int) T
}

func GeneratorIterator[T any](generateFunc func(int) T) commons.Iter[T] {
	if generateFunc == nil {
		panic(commons.ErrFuncIsNil)
	}
	return &generatorIterator[T]{generateFunc: generateFunc}
}

func (i *generatorIterator[T]) HasNext() bool {
	return true
}

func (i *generatorIterator[T]) Next() T {
	item := i.generateFunc(i.index)
	i.index += 1
	return item
}
