package bubblesort

import "errors"

func Sort(nums []int, rev bool) ([]int, error) {
	if len(nums) == 0 {
		return nil, errors.New("slice is empty")
	}

	rev_default := false
	if rev {
		rev_default = true
	}

	sorting := true

	for sorting {
		sorting = false
		for i := 1; i < len(nums); i++ {
			if rev_default {
				if nums[i-1] < nums[i] {
					sorting = true
					nums[i-1], nums[i] = nums[i], nums[i-1]
				}
			} else {
				if nums[i-1] > nums[i] {
					sorting = true
					nums[i-1], nums[i] = nums[i], nums[i-1]
				}
			}
		}
	}

	return nums, nil
}
