package _3_slides_windows

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 删掉一个元素以后全为 1 的最长子数组 https://leetcode.cn/problems/longest-subarray-of-1s-after-deleting-one-element/
func TestLongestSubarray1493(t *testing.T) {
	convey.Convey(" 每个字符最多出现两次的最长子字符串", t, func() {
		testCase := []struct {
			input    []int
			expected int
		}{
			{
				[]int{1, 1, 0, 1}, 3,
			},
			{
				[]int{0, 1, 1, 1, 0, 1, 1, 0, 1}, 5,
			},
			{
				[]int{1, 1, 1}, 2,
			},
		}

		for _, tst := range testCase {
			rsp := longestSubarray1493(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.expected)
		}
	})

}

func longestSubarray1493(nums []int) int {
	// 滑动窗口求最长

	var ans int
	var accur int
	var left int
	for right, num := range nums {
		if num == 0 {
			accur++ // 维护窗口中的 0 的个数
		}

		for accur > 1 {
			if nums[left] == 0 {
				accur--
			}
			left++
		}
		// 题目一定要删除一个数（尤其是全为 1 的情况下，要删除一个 1），所以删除后的长度是 right−left，
		ans = max(ans, right-left)
	}
	return ans
}
