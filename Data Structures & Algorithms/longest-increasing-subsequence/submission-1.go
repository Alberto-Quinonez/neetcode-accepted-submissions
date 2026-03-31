func lengthOfLIS(nums []int) int {
	res := 0
	dp := make([]int,len(nums))
	// loop all nums
	for i,_ := range nums {
		dp[i]=1
		// loop from 0 up until we reach before current outer loop location
		for j := 0; j < i; j++ {
			// scan all previous indices, check if I can extend current subsequence [i]
			// "can I extend the subsequence ending at j with nums[i]?
			// If yes, dp[i] = max(dp[i], dp[j] + 1).
			
			// check if prev is strictly less than current
			if nums[j] < nums[i] {
				// we assign dp[i] to max of either current dp[i] OR dp[j] + 1, 
				// The + 1 is due to fact that we are adding the previous subsequence to the newest addition to the sequence.
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}
