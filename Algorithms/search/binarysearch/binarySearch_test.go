package binarysearch

import "testing"

func TestSearchOne(t *testing.T) {
	nums := []int{1,2,3}
	ans := search(nums, 3)

	if ans != 2 {
		t.Errorf("Got %d, want 2", ans)
	}
}

func TestSearchTwo(t *testing.T) {
	nums := []int{1,3,4,5,10,13,17,32}
	ans := search(nums, 3)

	if ans != 1 {
		t.Errorf("Got %d, want 1", ans)
	}
}

func TestSearchThree(t *testing.T) {
	nums := []int{1,2,3}
	ans := search(nums, 4)

	if ans != -1 {
		t.Errorf("Got %d, want -1", ans)
	}
}