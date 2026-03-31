func dailyTemperatures(temperatures []int) []int {
	// init to 0 all indices
	res := make([]int,len(temperatures))
	mono := []int{}
	for i,t := range temperatures {
		// check if temperature is greater than
		for len(mono) > 0 &&  t > temperatures[mono[len(mono)-1]] {
			topIdx := mono[len(mono)-1]
			res[topIdx] = i - topIdx
			mono = mono[:len(mono)-1]
		}
		mono = append(mono, i)
	}
	return res
}
