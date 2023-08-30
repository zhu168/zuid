package zuid

import (
	"reflect"
	"sort"
	"testing"
)

func hasDuplicates(arr []string) bool {
	seen := make(map[string]bool) // Create a map to store seen elements.

	for _, item := range arr {
		if seen[item] {
			return true // If the element is already in the map, it's a duplicate.
		}
		seen[item] = true // Otherwise, mark the element as seen.
	}

	return false // If the loop completes without finding duplicates, return false.
}

func TestBigintIDGetID(t *testing.T) {
	zuid, err := NewZUID(1)
	if err != nil {
		t.Error(err.Error())
		t.Fail()
		return
	}
	var ls []string
	for i := 0; i < 5000; i++ {
		id := zuid.NextIDSimple()
		if len(id) != 32 {
			t.Error("zuid len err")
			t.Fail()
		}
		ls = append(ls, id)
	}
	newSlice := make([]string, len(ls))
	copy(newSlice, ls)
	sort.Strings(newSlice)
	if !reflect.DeepEqual(ls, newSlice) {
		t.Error("String slices are not equal.")
		t.Fail()
	}
	if hasDuplicates(ls) {
		t.Error("zuid contains duplicate elements.")
		t.Fail()
	}
}
