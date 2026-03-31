func minWindow(s string, t string) string {
	// frequency map
	freq := map[rune]int{}
	for _,c := range t {
		freq[c]++
	}
	// we will retieve substring with index found after
	smallestIdx := []int{0,0}
	// no guarantee that we find a smallest window. 
	// if none found return empty string
	smallestWin := math.MaxInt
	// init since we need all runes initially
	need := len(freq)
	formed := 0
	left := 0
	right := 0
	
	// start moving pointers
	// stop once we hit end
	sArr := []rune(s)
	for right < len(s) {
		// check if character in t
		if _, ok := freq[sArr[right]]; ok {
			freq[sArr[right]]--
			if freq[sArr[right]] == 0 {
				formed++
			}
		}
		// check for valid substring
		for formed == need {
			// if valid, check if this window is smaller than current smallest
			currWin := right - left + 1
			if currWin <  smallestWin {
				// update smallest and index
				smallestWin = currWin
				smallestIdx[0] = left
				smallestIdx[1] = right + 1
			}
			// we also want to constrict the left side
			if _, ok := freq[sArr[left]]; ok {
				freq[sArr[left]]++
				if freq[sArr[left]] > 0 {
					formed--
				}
			}
			left++
		}
		right++
	}
	if smallestWin == math.MaxInt {
		return ""	
	} 
	//once we exit, build up the substring with
	substr := sArr[smallestIdx[0]:smallestIdx[1]]
	return string(substr)
}
