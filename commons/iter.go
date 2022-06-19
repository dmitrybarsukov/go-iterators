package commons

type Iter[T any] interface {
	HasNext() bool
	Next() T
}
