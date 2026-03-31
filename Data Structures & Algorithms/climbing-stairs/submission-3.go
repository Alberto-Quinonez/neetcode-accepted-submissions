func climbStairs(n int) int {
	c := map[int]int{}
	return mcs(n,c)
}

func mcs(n int, c map[int]int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if c[n] != 0 {
		return c[n]
	}
	res := mcs(n-1, c) + mcs(n-2, c)
	c[n] = res	
	return res
}