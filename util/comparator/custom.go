package comparator

import (
	"reflect"
	"time"

	"iterator/commons"
)

func compare[T commons.Ordered](i, j any) bool {
	return i.(T) < j.(T)
}

var typeToComparatorMap = map[reflect.Type]func(any, any) bool{
	reflect.TypeOf(int(0)):           compare[int],
	reflect.TypeOf(int8(0)):          compare[int8],
	reflect.TypeOf(int16(0)):         compare[int16],
	reflect.TypeOf(int32(0)):         compare[int32],
	reflect.TypeOf(int64(0)):         compare[int64],
	reflect.TypeOf(uint(0)):          compare[uint],
	reflect.TypeOf(uint8(0)):         compare[uint8],
	reflect.TypeOf(uint16(0)):        compare[uint16],
	reflect.TypeOf(uint32(0)):        compare[uint32],
	reflect.TypeOf(uint64(0)):        compare[uint64],
	reflect.TypeOf(float32(0)):       compare[float32],
	reflect.TypeOf(float64(0)):       compare[float64],
	reflect.TypeOf(string("")):       compare[string],
	reflect.TypeOf(time.Duration(0)): compare[time.Duration],
}

func init() {
	RegisterComparator(func(t1, t2 time.Time) bool { return t1.Before(t2) })
}

func RegisterComparator[T any](compFunc func(T, T) bool) {
	var t T
	typeToComparatorMap[reflect.TypeOf(t)] = func(i, j any) bool {
		it, _ := i.(T)
		jt, _ := j.(T)
		return compFunc(it, jt)
	}
}

func getComparator(reflectType reflect.Type) func(any, any) bool {
	if comp, ok := typeToComparatorMap[reflectType]; ok {
		return comp
	} else {
		panic(commons.ErrTypeIsNotComparable)
	}
}
