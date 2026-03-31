func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
    	return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{}
	for i := 0; i < len(intervals); i++ {
		// check for empty result, append immediately
		// otherwise check if previous result's last val is less than currents first.
		// if true, nothing to merge
		if len(res) == 0 || res[len(res)-1][1] < intervals[i][0] {
			res = append(res, intervals[i])
		} else {
			// merge, set previous to maximum of previous end and new end. 
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
		}
	}
	return res
}
