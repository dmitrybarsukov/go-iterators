package iterator

type Iter[T any] interface {
	HasNext() bool
	Next() T
}

type Iterator[T any] struct {
	innerIter Iter[T]
}

func (i *Iterator[T]) HasNext() bool {
	return i.innerIter.HasNext()
}

func (i *Iterator[T]) Next() T {
	return i.innerIter.Next()
}
