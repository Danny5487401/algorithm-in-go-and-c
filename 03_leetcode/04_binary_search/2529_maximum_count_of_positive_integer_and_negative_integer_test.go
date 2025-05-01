package _4_binary_search

import (
	"github.com/smartystreets/goconvey/convey"
	"sort"
	"testing"
)

// 正整数和负整数的最大计数 https://leetcode.cn/problems/maximum-count-of-positive-integer-and-negative-integer/description/
func TestMaximumCount(t *testing.T) {
	convey.Convey("按 非递减顺序 排列的数组 nums,正整数和负整数的最大计数,如果 nums 中正整数的数目是 pos ，而负整数的数目是 neg ，返回 pos 和 neg二者中的最大值", t, func() {
		testCase := []struct {
			input    []int
			expected int
		}{
			{

				[]int{-2, -1, -1, 1, 2, 3}, 3,
			},
			{
				[]int{-3, -2, -1, 0, 0, 1, 2}, 3,
			},
			{
				[]int{5, 20, 66, 1314}, 4,
			},
		}

		for _, tst := range testCase {
			rsp := maximumCount(tst.input)
			convey.So(rsp, convey.ShouldResemble, tst.expected)
		}
	})

}

func maximumCount(nums []int) int {
	var neg, pos int
	length := len(nums) - 1
	// 二分找到第一个 ≥0 的数的下标 i，那么下标在 [0,i−1] 中的数都小于 0，这恰好有 i 个
	// 二分找到第一个 >0 的数的下标 j，那么下标在 [j,n−1] 中的数都大于 0，这有 n−j 个
	left := lowerBound3(nums, 0)
	neg = left

	right := lowerBound3(nums, 1)

	pos = length - right + 1

	return max(neg, pos)

}

func maximumCount2(nums []int) int {
	neg := sort.SearchInts(nums, 0)
	// 第一个 > 0 的位置，等价于第一个 >= 1 的位置
	pos := len(nums) - sort.SearchInts(nums, 1)
	return max(neg, pos)
}
