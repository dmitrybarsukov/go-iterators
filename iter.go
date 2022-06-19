package go_iterators

type Iter[T any] interface {
	HasNext() bool
	Next() T
}
