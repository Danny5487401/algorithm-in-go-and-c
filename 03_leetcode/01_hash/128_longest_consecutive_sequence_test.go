package main

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 最长连续序列 https://leetcode.cn/problems/longest-consecutive-sequence/description/?envType=study-plan-v2&envId=top-100-liked
func TestLongestConsecutive(t *testing.T) {
	convey.Convey("longestConsecutive ", t, func() {
		testCase := []struct {
			input    []int
			expected int
		}{
			{
				[]int{100, 4, 200, 1, 3, 2}, 4,
			},
		}

		for _, tst := range testCase {
			ret := longestConsecutive(tst.input)
			convey.So(ret, convey.ShouldEqual, tst.expected)
		}
	})

}

func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}
	// 最长数字连续序列数
	longestStreak := 0
	for _, num := range nums {
		// 如果已知有一个 x,x+1,x+2,⋯,x+y 的连续序列，而我们却重新从 x+1，x+2 或者是 x+y处开始尝试匹配，那么得到的结果肯定不会优于枚举 x 为起点的答案
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if longestStreak < currentStreak {
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak
}
