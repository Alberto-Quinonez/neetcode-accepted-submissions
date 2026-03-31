func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		if i> 0 && nums[i] == nums[i-1] {
			continue
		}
		// init ptrs
		left,right := i+1, len(nums)-1
		for left < right {
			localSum := nums[i] + nums[left] + nums[right]
			if localSum == 0 {
            	res = append(res, []int{nums[i], nums[left], nums[right]})
              	left++
            	for left < right && nums[left] == nums[left-1] { 
					left++ 
				}
			} else if localSum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}
