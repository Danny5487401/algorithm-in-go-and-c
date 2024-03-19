package main

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 盛最多水的容器 https://leetcode.cn/problems/container-with-most-water/description/?envType=study-plan-v2&envId=top-100-liked

func TestMaxArea(t *testing.T) {
	convey.Convey("Container With Most Water ", t, func() {
		testCase := []struct {
			input    []int
			expected int
		}{
			{
				[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49,
			},
			{
				[]int{1, 1}, 1,
			},
		}

		for _, tst := range testCase {
			rsp := maxArea(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.expected)
		}
	})

}

func maxArea(height []int) int {

	// 两头往里走时，长度肯定变短。如果高度变矮了，肯定面积变小。所以应该选择高度变大
	var ans = 0
	left, right := 0, len(height)-1
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		ans = max(ans, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}
