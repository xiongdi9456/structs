package structs

import (
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {
	type S struct {
		A string
		B int
		C bool
	}
	T := []S{
		{"a", 2, true},
		{"b", 3, false},
	}

	a := SliceMap(T)

	if len(a) != 2 {
		t.Errorf("SliceMap should contain a slice len 2, got %d", len(a))
	}

	inMap := func(key interface{}, idx int) bool {
		for _, v := range a[idx] {
			if reflect.DeepEqual(v, key) {
				return true
			}
		}
		return false
	}

	for i, ele := range [][]interface{}{{"a", 2, true}, {"b", 3, false}} {
		for _, val := range ele {
			if !inMap(val, i) {
				t.Errorf("Map should have the value %v", val)
			}
		}
	}
}
