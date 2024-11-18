package _3_slides_windows

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 绝对差不超过限制的最长连续子数组 https://leetcode.cn/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/description/
func TestLongestSubarray(t *testing.T) {
	convey.Convey("绝对差不超过限制的最长连续子数组:任意两个元素之间的绝对差必须小于或者等于 limit", t, func() {
		testCase := []struct {
			nums     []int
			limit    int
			expected int
		}{
			{
				[]int{8, 2, 4, 7}, 4, 2,
			},
			{
				[]int{10, 1, 2, 4, 7, 2}, 5, 4,
			},
			{
				[]int{4, 2, 2, 2, 4, 4, 2, 2}, 0, 3,
			},
		}

		for _, tst := range testCase {
			rsp := longestSubarray(tst.nums, tst.limit)
			convey.So(rsp, convey.ShouldEqual, tst.expected)
		}
	})

}

func longestSubarray(nums []int, limit int) (ans int) {
	// 整数 1 <= nums[i] <= 10^9

	// 使用一个单调递增的队列 queMin 维护最小值，一个单调递减的队列 queMax 维护最大值。
	// 这样我们只需要计算两个队列的队首的差值
	var minQ, maxQ []int
	left := 0
	for right, v := range nums {
		for len(minQ) > 0 && minQ[len(minQ)-1] > v {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, v)

		for len(maxQ) > 0 && maxQ[len(maxQ)-1] < v {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, v)

		for len(minQ) > 0 && len(maxQ) > 0 && maxQ[0]-minQ[0] > limit { // 如果之间的差大于limit，那就需要改变最大or最小
			// 比较单调栈的首元素，因为最大的和最小的都是队首
			if nums[left] == minQ[0] {
				minQ = minQ[1:]
			}
			if nums[left] == maxQ[0] {
				maxQ = maxQ[1:]
			}
			left++
		}
		ans = max(ans, right-left+1) // 返回的是长度
	}
	return
}
