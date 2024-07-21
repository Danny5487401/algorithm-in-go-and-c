package _4_binary_search

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 寻找旋转排序数组中的最小值 https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/solution/by-endlesscheng-owgd/
func TestFindMin(t *testing.T) {
	convey.Convey("寻找旋转排序数组中的最小值 ", t, func() {
		testCase := []struct {
			input    []int
			expected int
		}{
			{

				[]int{3, 4, 5, 1, 2}, 1,
			},
			{
				[]int{4, 5, 6, 7, 0, 1, 2}, 0,
			},
			{
				[]int{11, 13, 15, 17}, 11,
			},
		}

		for _, tst := range testCase {
			rsp := findMin(tst.input)
			convey.So(rsp, convey.ShouldResemble, tst.expected)
		}
	})

}

func findMin(nums []int) int {
	// 二分的范围 [0,n-2] --> 开区间 (-1,n-1)
	var left = -1
	var right = len(nums) - 1

	for left+1 < right {
		mid := left + (right-left)/2

		if nums[mid] < nums[len(nums)-1] {
			right = mid
		} else {
			left = mid
		}

	}

	return nums[right]
}
