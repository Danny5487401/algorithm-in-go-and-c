package _4_binary_search

import (
	"github.com/smartystreets/goconvey/convey"
	"sort"
	"testing"
)

// 正整数和负整数的最大计数 https://leetcode.cn/problems/maximum-count-of-positive-integer-and-negative-integer/description/
func TestMaximumCount(t *testing.T) {
	convey.Convey("按 非递减顺序 排列的数组 nums,正整数和负整数的最大计数,"+
		"如果 nums 中正整数的数目是 pos ，而负整数的数目是 neg ，返回 pos 和 neg二者中的最大值, 0 既不是正整数也不是负整数", t, func() {
		testCase := []struct {
			input    []int
			expected int
		}{
			{
				// 共有 3 个正整数和 3 个负整数。计数得到的最大值是 3 。
				[]int{-2, -1, -1, 1, 2, 3}, 3,
			},
			{
				// 共有 2 个正整数和 3 个负整数。计数得到的最大值是 3 。
				[]int{-3, -2, -1, 0, 0, 1, 2}, 3,
			},
			{
				// 共有 4 个正整数和 0 个负整数。计数得到的最大值是 4
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
	// 负数: 二分找到第一个 <0 --> ≥0-1
	// 正数: 二分找到第一个 >0 --> >=1
	left := lowerBound3(nums, 0) - 1

	neg = left + 1 // index +1 代表数量

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
