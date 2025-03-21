package aign

func MaxArea(height []int) int {
	left, right := 0, len(height)-1
	var area = 0
	for left <= right {
		area = max(area, min(height[left], height[right])*(right-left))
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return area
}
