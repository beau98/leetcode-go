package top150

import "math"

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	res := 0
	area := 0
	for l < r {
		if height[l] > height[r] {
			area = (r - l) * height[r]
			r--
		} else {
			area = (r - l) * height[l]
			l++
		}
		res = int(math.Max(float64(res), float64(area)))
	}
	return res
}
