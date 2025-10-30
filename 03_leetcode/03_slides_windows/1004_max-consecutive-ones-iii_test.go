package _3_slides_windows

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 最大连续1的个数 III  https://leetcode.cn/problems/max-consecutive-ones-iii/description/
func TestLongestOnes(t *testing.T) {
	convey.Convey("给定一个二进制数组 nums 和一个整数 k，假设最多可以翻转 k 个 0 ，则返回执行操作后 数组中连续 1 的最大个数 ", t, func() {
		testCase := []struct {
			nums     []int
			k        int
			expected int
		}{
			{
				// 子数组 [4,3]
				[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6,
			},
			{
				[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3, 10,
			},
		}

		for _, tst := range testCase {
			rsp := longestOnes(tst.nums, tst.k)
			convey.So(rsp, convey.ShouldEqual, tst.expected)
		}
	})

}

func longestOnes(nums []int, k int) int {

	zeroCount := 0
	ans := 0
	left := 0
	for right, num := range nums {
		if num == 0 {
			zeroCount++
		}
		for zeroCount > k {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
