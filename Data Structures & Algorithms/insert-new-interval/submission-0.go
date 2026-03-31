func insert(intervals [][]int, newInterval []int) [][]int {
    res := [][]int{}
	for _, interval := range intervals {
		// check the interval
		// 1st pass, appending all intervals that are stricly less than new interval 
		if interval[1] < newInterval[0] {
			res = append(res, interval)
			//handle case when we are past the newInterval
		} else if interval[0] > newInterval[1]{
			res = append(res, newInterval)
			// reseting new Interval to keep appending it.
			newInterval = interval
			// merging, not ready to append yet.
		} else {
			newInterval[0] = min(newInterval[0], interval[0])
			newInterval[1] = max(newInterval[1], interval[1])
		}
	}
	res = append(res, newInterval)
	return res
}
