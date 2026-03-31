func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := 0;i<len(dp);i++{
		dp[i] = make([]int, len(text2)+1)
	}
	// outer loop, do not go full length, extra is to hold 0 base cases
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			// check for matching text
			if text1[i-1] == text2[j-1] {
				// fetch diagonal value + 1
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// max of bottom or right.
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}


    // table
    //             a    c    e
    //        0    0    0    0
    //   a    0    1    1    1
    //   b    0    1    1    1
    //   c    0    1    2    2
    //   d    0    1    2    2
    //   e    0    1    2    3