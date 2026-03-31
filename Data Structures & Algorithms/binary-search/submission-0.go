func search(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		// found it
		if nums[mid] == target {
			return mid
		}
		// target is on left side
		// bring right to mids spot - 1
		if nums[mid] > target {
			right = mid-1
		}
		// target is on right side
		// bring right to mids spot + 1
		if nums[mid] < target {
			left = mid+1
		}
	}
	return -1
}
