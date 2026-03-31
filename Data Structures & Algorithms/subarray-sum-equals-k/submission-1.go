func subarraySum(nums []int, k int) int {
	// map containing cummulative sum at each index
	prefixCounts := map[int]int{}
	prefixCounts[0] = 1
	runningSum := 0
	res := 0
	for _,n := range nums {
		runningSum += n
		// look for subarrays that 
		if val,ok := prefixCounts[runningSum-k]; ok {
			res += val
		}
		prefixCounts[runningSum]++
	}
	return res
}
