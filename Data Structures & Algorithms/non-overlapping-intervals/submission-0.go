func eraseOverlapIntervals(intervals [][]int) int {	
	// sort by end time
	sort.Slice(intervals, func(a,b int) bool{
		return intervals[a][1] < intervals[b][1]
	})
	skipped := 0
	// initialize lastKeptEnd with end of 1st interval
	lastKeptEnd := intervals[0][1]
	// start at 2nd interval
	for i:=1;i<len(intervals);i++ {
		if intervals[i][0] < lastKeptEnd{
			// overlap found, skipped
			skipped++
		} else {
			lastKeptEnd = intervals[i][1]
		}
	}
	return skipped
}
