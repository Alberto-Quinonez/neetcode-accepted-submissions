func maxArea(heights []int) int {
	left, right := 0,len(heights)-1
	maxArea := 0
	for left < right {
		area := min(heights[left],heights[right]) * (right-left)
		if area > maxArea {
			maxArea = area
		}
		if heights[left] > heights[right] {
			right--
		} else {
			left++
		}
	}
	return maxArea
}
