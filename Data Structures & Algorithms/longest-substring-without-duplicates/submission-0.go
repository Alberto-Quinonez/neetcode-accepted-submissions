func lengthOfLongestSubstring(s string) int {
	// map to hold location of last seen location of character
	charSeen := map[rune]int{}
	res := 0
	left := 0
	right := 0
	for _,c := range s {
		// fetch last seen location for char
		val, ok := charSeen[c]
		if ok && val >= left {
			// duplicate found
			left = val + 1
		}
		//update char map
		charSeen[c] = right
		// move right ptr
		right++
		res = max(res,right-left)
	}
	return res
}

//trace
//Input: s = "zxyzxyz"
//Output: 3 

// loop first rune z
// z not in map
// currWinSize = 1
// charSeen[`z`] = 0
// right = 1
// res = max(0,1) = 1

// next rune x
// x not in map
// currWinSize = 2
// charSeen[`x`] = 1
// right = 2
// res = max(1,2) = 2

// next rune y
// y not in map
// currWinSize = 3
// charSeen[`y`] = 2
// right = 3
// res = max(2,3) = 3

// next rune z
// z is in map
// val is 0, 0 <= right && 0>= left (left == 0)
// currWinSize = 0
// left = val+1 = 0+1 = 1
// charSeen[`z`] = 3
// right = 4
// res = max(3,0) = 3





