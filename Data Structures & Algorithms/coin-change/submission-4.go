func coinChange(coins []int, amount int) int {
	cache := map[int]int{}
	return memoCC(coins, amount, cache)
}

func memoCC(coins []int, amount int, cache map[int]int) int {
	if amount == 0 {
		return 0
	}
	res := amount + 1
	for _, c := range coins {
		if amount - c >= 0 {
			var sub int
			if val, ok := cache[amount-c]; ok {
				sub = val
			} else {
				sub = memoCC(coins, amount-c, cache)
				cache[amount-c] = sub
			}
			if sub != -1 {
				res = min(res, 1+sub)
			}
		}
	}
	if res > amount {
		return -1
	}
	return res
}