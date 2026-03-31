func subsets(nums []int) [][]int {
	res := [][]int{}
	path := []int{}
	var bt func(start int)
	bt = func(start int) {
		res = append(res, append([]int{}, path...))
		for i := start;i < len(nums); i++ {
			// building path for given i
			path = append(path,nums[i])
			bt(i+1)
			path = path[:len(path)-1]
		}
	}
	bt(0)
	return res
}