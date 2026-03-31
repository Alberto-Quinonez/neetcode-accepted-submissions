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

// nums=[5,1,3]
// target=5

// left = 0, right = 2

// mid = 1
// nums[1] == 1 != target
// nums[0] <= nums[1] == false -> left sorted
// nums[1] <= target && target <= nums[2] -> false. target 5 is greater than nums[2] = 3
// right = 0
// mid = (0+0)/2 = 0
// nums[0] = 5 == target true