func search(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	for left <= right {
		mid := (left+right)/2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] {
			// left side sorted
			if nums[left] <= target && target <= nums[mid] {
				right = mid-1
			} else {
				left = mid+1
			}
		} else {
			// right side sorted
			if nums[mid] <= target && target <= nums[right] {
				left = mid+1
			} else {
				right = mid-1
			}
		}
	}
	return -1
}