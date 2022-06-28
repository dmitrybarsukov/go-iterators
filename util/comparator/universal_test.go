package comparator

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUniversalComparator(t *testing.T) {
	testCases := []struct {
		name string
		arg1 any
		arg2 any
	}{
		{
			name: "int",
			arg1: 1,
			arg2: 2,
		},
		{
			name: "uint",
			arg1: uint(1),
			arg2: uint(2),
		},
		{
			name: "string",
			arg1: "hello",
			arg2: "world",
		},
		{
			name: "Duration",
			arg1: time.Duration(0),
			arg2: time.Duration(5),
		},
		{
			name: "Time",
			arg1: time.Time{},
			arg2: time.Now(),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			comp := NewUniversalComparator()
			assert.True(t, comp.Compare(tc.arg1, tc.arg2))
			assert.False(t, comp.Compare(tc.arg2, tc.arg1))
		})
	}
}

func TestRegisterCustomComparator(t *testing.T) {
	RegisterComparator(func(i1, i2 *int) bool {
		if i1 == nil {
			return true
		} else if i2 == nil {
			return false
		} else {
			return *i1 < *i2
		}
	})
	one := 1
	two := 2
	assert.True(t, NewUniversalComparator().Compare(nil, &one))
	assert.False(t, NewUniversalComparator().Compare(&one, nil))
	assert.False(t, NewUniversalComparator().Compare(nil, nil))
	assert.True(t, NewUniversalComparator().Compare(&one, &two))
	assert.False(t, NewUniversalComparator().Compare(&two, &one))
}
