package _4_binary_search

import (
	"github.com/smartystreets/goconvey/convey"
	"sort"
	"testing"
)

//  在排序数组中查找元素的第一个和最后一个位置 https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/description/?envType=study-plan-v2&envId=top-100-liked

func TestSearchRange(t *testing.T) {
	convey.Convey("Move Zeroes ", t, func() {
		testCase := []struct {
			input    []int
			target   int
			expected []int
		}{
			{

				[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1. - 1},
			},
			{
				[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4},
			},
		}

		for _, tst := range testCase {
			rsp := searchRange(tst.input, tst.target)
			convey.So(rsp, convey.ShouldEqual, tst.expected)
		}
	})

}

func searchRange(nums []int, target int) []int {
	// 这里必须是升序
	leftMost := sort.SearchInts(nums, target)
	if leftMost < len(nums) && nums[leftMost] == target {
		rightmost := sort.SearchInts(nums, target+1) - 1
		return []int{leftMost, rightmost}
	} else {
		return []int{-1. - 1}
	}
}
