package util

import (
	"errors"
	"testing"

	"iterator/basic"

	"github.com/stretchr/testify/assert"
)

func TestToSlice(t *testing.T) {
	items := []int{1, 2, 3, 4}
	iter := basic.SliceIterator(items)
	assert.Equal(t, items, ToSlice(iter))
}

func TestToMap(t *testing.T) {
	items := []int{1, 2, 3, 4}
	iter := basic.SliceIterator(items)
	actual := ToMap(iter, func(i int) int { return i - 1 })
	expected := map[int]int{0: 1, 1: 2, 2: 3, 3: 4}
	assert.Equal(t, expected, actual)
}

func TestToMapKeyValue(t *testing.T) {
	items := map[int]int{1: 2, 2: 4, 4: 8}
	iter := basic.MapIterator(items)
	actual := ToMapKeyValue(iter)
	assert.Equal(t, items, actual)
}

func TestToMapWithValue(t *testing.T) {
	items := []int{1, 2, 3, 4}
	iter := basic.SliceIterator(items)
	actual := ToMapWithValue(iter, func(i int) int { return i - 1 }, func(i int) int { return i + 1 })
	expected := map[int]int{0: 2, 1: 3, 2: 4, 3: 5}
	assert.Equal(t, expected, actual)
}

func TestCount(t *testing.T) {
	items := []int{1, 2, 3, 4}
	iter := basic.SliceIterator(items)
	assert.Equal(t, 4, Count(iter))
}

func TestCountMatching(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	iter := basic.SliceIterator(items)
	assert.Equal(t, 3, CountMatching(iter, func(i int) bool { return i%2 == 1 }))
}

func TestFirst(t *testing.T) {
	_, ok := First(basic.SliceIterator([]int{}))
	assert.False(t, ok)

	value, ok := First(basic.SliceIterator([]int{1, 2, 3}))
	assert.True(t, ok)
	assert.Equal(t, 1, value)
}

func TestFirstOrZeroValue(t *testing.T) {
	value := FirstOrZeroValue(basic.SliceIterator([]int{}))
	assert.Zero(t, value)
}

func TestFirstOrDefault(t *testing.T) {
	value := FirstOrDefault(basic.SliceIterator([]int{}), 5)
	assert.Equal(t, 5, value)
}

func TestLast(t *testing.T) {
	_, ok := Last(basic.SliceIterator([]int{}))
	assert.False(t, ok)

	value, ok := Last(basic.SliceIterator([]int{1, 2, 3}))
	assert.True(t, ok)
	assert.Equal(t, 3, value)
}

func TestLastOrZeroValue(t *testing.T) {
	value := LastOrZeroValue(basic.SliceIterator([]int{}))
	assert.Zero(t, value)
}

func TestLastOrDefault(t *testing.T) {
	value := LastOrDefault(basic.SliceIterator([]int{}), 5)
	assert.Equal(t, 5, value)
}

func TestMaxBy(t *testing.T) {
	_, ok := MaxBy(basic.SliceIterator([]int{}), Self[int])
	assert.False(t, ok)

	value, ok := MaxBy(basic.SliceIterator([]int{1, 3, 2}), Self[int])
	assert.True(t, ok)
	assert.Equal(t, 3, value)
}

func TestMaxByOrZeroValue(t *testing.T) {
	value := MaxByOrZeroValue(basic.SliceIterator([]int{}), Self[int])
	assert.Zero(t, value)
}

func TestMaxByOrDefault(t *testing.T) {
	value := MaxByOrDefault(basic.SliceIterator([]int{}), Self[int], 8)
	assert.Equal(t, 8, value)
}

func TestMinBy(t *testing.T) {
	_, ok := MinBy(basic.SliceIterator([]int{}), Self[int])
	assert.False(t, ok)

	value, ok := MinBy(basic.SliceIterator([]int{1, 3, 2}), Self[int])
	assert.True(t, ok)
	assert.Equal(t, 1, value)
}

func TestMinByOrZeroValue(t *testing.T) {
	value := MinByOrZeroValue(basic.SliceIterator([]int{}), Self[int])
	assert.Zero(t, value)
}

func TestMinByOrDefault(t *testing.T) {
	value := MinByOrDefault(basic.SliceIterator([]int{}), Self[int], 8)
	assert.Equal(t, 8, value)
}

func TestForEach(t *testing.T) {
	items := []int{1, 2, 3}
	var result []int
	ForEach(basic.SliceIterator(items), func(i int) { result = append(result, i) })
	assert.Equal(t, items, result)
}

func TestForEachCollectingErrors(t *testing.T) {
	items := []string{"1", "2", "3"}
	errs := ForEachCollectingErrors(basic.SliceIterator(items), func(i string) error {
		return errors.New(i)
	})
	assert.Equal(t, []error{errors.New("1"), errors.New("2"), errors.New("3")}, errs)
}

func TestForEachUntilFirstError(t *testing.T) {
	err := ForEachUntilFirstError(basic.SliceIterator([]int{1, 2, 3}), func(i int) error {
		if i == 2 {
			return errors.New("2")
		} else {
			return nil
		}
	})
	assert.Equal(t, errors.New("2"), err)

	err = ForEachUntilFirstError(basic.SliceIterator([]int{1, 2, 3}), func(i int) error { return nil })
	assert.Equal(t, nil, err)
}

func TestAllAnyNone(t *testing.T) {
	isEven := func(i int) bool { return i%2 == 0 }
	testCases := []struct {
		name         string
		items        []int
		expectedAll  bool
		expectedAny  bool
		expectedNone bool
	}{
		{
			name:         "None is even",
			items:        []int{1, 3, 5},
			expectedAll:  false,
			expectedAny:  false,
			expectedNone: true,
		},
		{
			name:         "Some is even",
			items:        []int{1, 2, 3},
			expectedAll:  false,
			expectedAny:  true,
			expectedNone: false,
		},
		{
			name:         "All is even",
			items:        []int{2, 4, 6},
			expectedAll:  true,
			expectedAny:  true,
			expectedNone: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedAll, All(basic.SliceIterator(tc.items), isEven))
			assert.Equal(t, tc.expectedAny, Any(basic.SliceIterator(tc.items), isEven))
			assert.Equal(t, tc.expectedNone, None(basic.SliceIterator(tc.items), isEven))
		})
	}
}
