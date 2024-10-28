package _3_slides_windows

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 最多 K 个重复元素的最长子数组 https://leetcode.cn/problems/length-of-longest-subarray-with-at-most-k-frequency/
func TestMaxSubarrayLength(t *testing.T) {
	convey.Convey("最多 K 个重复元素的最长子数组", t, func() {
		testCase := []struct {
			input    []int
			k        int
			expected int
		}{
			{
				[]int{1, 2, 3, 1, 2, 3, 1, 2}, 2, 6,
			},
			{
				[]int{1, 2, 1, 2, 1, 2, 1, 2}, 1, 2,
			},
		}

		for _, tst := range testCase {
			rsp := maxSubarrayLength(tst.input, tst.k)
			convey.So(rsp, convey.ShouldEqual, tst.expected)
		}
	})

}

func maxSubarrayLength(nums []int, k int) int {
	var count = make(map[int]int) // 计算出现次数

	left := 0
	length := len(nums)
	if length <= 1 {
		return length
	}

	var ans int

	for right, num := range nums {
		count[num]++
		for count[num] > k { // 无重复字符的最长子串 https://leetcode.cn/problems/longest-substring-without-repeating-characters/solution/xia-biao-zong-suan-cuo-qing-kan-zhe-by-e-iaks/
			count[nums[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans

}
