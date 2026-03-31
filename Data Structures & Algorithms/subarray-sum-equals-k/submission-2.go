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


// prefixCounts[0] = 1
// runningSum := 0
// res := 0

// first loop
// n = 1
// runningSum = 0 + 1 = 1
// prefixCounts[1-3 = -2] not found
// prefixCounts[1] = 1

// 2nd loop
// n = 2
// runningSum = 1 + 2 = 3
// prefixCounts[3-3 = 1] = 1
// res = 1
// prefixCounts[3] = 1

// 3rd loop
// n = 3
// runningSum = 3 + 3 = 6
// prefixCounts[6-3 = 3] = 1
// res = 1+1 = 2
// prefixCounts[6] = 1

// return res = 2