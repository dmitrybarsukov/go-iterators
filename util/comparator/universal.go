package comparator

import "reflect"

type UniversalComparator struct {
	innerComparator func(any, any) bool
}

func NewUniversalComparator() *UniversalComparator {
	return &UniversalComparator{}
}

func (c *UniversalComparator) Compare(i, j any) bool {
	if c.innerComparator == nil {
		if i != nil {
			c.innerComparator = getComparator(reflect.TypeOf(i))
		} else if j != nil {
			c.innerComparator = getComparator(reflect.TypeOf(j))
		} else {
			return false
		}
	}
	return c.innerComparator(i, j)
}
