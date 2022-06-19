package commons

type KeyValue[TK comparable, TV any] struct {
	Key   TK
	Value TV
}

func (v KeyValue[TK, TV]) Pair() (TK, TV) {
	return v.Key, v.Value
}
