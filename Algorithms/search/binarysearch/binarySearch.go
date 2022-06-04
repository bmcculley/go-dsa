package binarysearch

func search(nums []int, key int) int {
	left := 0
	right := len(nums) - 1
	var mid int

	for left <= right {
		mid = (left + right) / 2
		
		if (nums[mid] == key) {
			return mid
		} else if (nums[mid] < key) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}