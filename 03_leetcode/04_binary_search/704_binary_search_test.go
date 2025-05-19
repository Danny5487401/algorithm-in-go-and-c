package _4_binary_search

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 二分查找 https://leetcode.cn/problems/binary-search/description/
func TestSearch2(t *testing.T) {
	convey.Convey("二分查找: n 个元素有序的（升序）整型数组,如果目标值存在返回下标，否则返回 -1", t, func() {
		testCase := []struct {
			input    []int
			target   int
			expected int
		}{
			{
				[]int{-1, 0, 3, 5, 9, 12}, 9, 4,
			},
			{
				[]int{-1, 0, 3, 5, 9, 12}, 2, -1,
			},
		}

		for _, tst := range testCase {
			rsp := search2(tst.input, tst.target)
			convey.So(rsp, convey.ShouldResemble, tst.expected)
		}
	})

}

func search2(nums []int, target int) int {
	// 求大于等于
	index := lowerBound2(nums, target)
	if index == len(nums) || nums[index] != target {
		return -1
	}
	return index
}
