func coinChange(coins []int, amount int) int {
	if amount <= 0 {
		return 0
	}
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
		for _, c := range coins {
			if c <= i {
				dp[i] = min(dp[i], 1 + dp[i-c])
			}
			
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}