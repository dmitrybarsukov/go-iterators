package basic

import "iterator/commons"

type intRangeIterator struct {
	startInclusive int
	endExclusive   int
	step           int
	current        int
}

func IntRangeIterator(startInclusive, endExclusive, step int) commons.Iter[int] {
	if step == 0 {
		panic(commons.ErrArgumentStepIsZero)
	}
	if endExclusive-startInclusive > 0 && step < 0 || endExclusive-startInclusive < 0 && step > 0 {
		panic(commons.ErrArgumentStepHasWrongSign)
	}
	return &intRangeIterator{
		startInclusive: startInclusive,
		endExclusive:   endExclusive,
		step:           step,
		current:        startInclusive,
	}
}

func (i *intRangeIterator) HasNext() bool {
	if i.step > 0 {
		return i.current < i.endExclusive
	} else {
		return i.current > i.endExclusive
	}
}

func (i *intRangeIterator) Next() int {
	result := i.current
	i.current += i.step
	return result
}
