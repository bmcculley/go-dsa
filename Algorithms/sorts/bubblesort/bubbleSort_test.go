package bubblesort

import "testing"

func compareSlice(sorted, expected []int) bool {
	if len(sorted) != len(expected) {
		return false
	}
	for i, v := range sorted {
		if v != expected[i] {
			return false
		}
	}
	return true
}

func TestSortOne(t *testing.T) {
	nums := []int{2, 1, 9, 4, 3}
	
	expected := []int{1, 2, 3, 4, 9}

	nums, _ = Sort(nums, false)

	if (!compareSlice(nums, expected)) {
		t.Errorf("Got %q, want %q", nums, expected)
	}
}

func TestSortTwo(t *testing.T) {
	nums := []int{2, 1, 9, 4, 3}
	
	expected := []int{9, 4, 3, 2, 1}

	nums, _ = Sort(nums, true)

	if (!compareSlice(nums, expected)) {
		t.Errorf("Got %q, want %q", nums, expected)
	}
}