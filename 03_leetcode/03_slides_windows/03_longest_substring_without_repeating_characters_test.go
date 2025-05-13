package _3_slides_windows

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 无重复字符的最长子串 https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/?envType=study-plan-v2&envId=top-100-liked

func TestLengthOfLongestSubstring(t *testing.T) {
	convey.Convey("无重复字符的最长子串", t, func() {
		testCase := []struct {
			input    string
			expected int
		}{
			{
				"abcabcbb", 3,
			},
			{
				"bbbbb", 1,
			},
			{
				"pwwkew", 3,
			},
		}

		for _, tst := range testCase {
			rsp := lengthOfLongestSubstring(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.expected)
		}
	})

}

func lengthOfLongestSubstring(s string) int {
	// 假设第k个字符作为起始位置，最长结束位置为Rk,那么第k+1到rk的字符肯定也是不重复，所以只要移动rk直到重复字符。
	length := len(s)
	if length <= 1 {
		return length
	}
	ans := 0
	// 哈希判断是否有重复的字符, key为数据 value出现次数
	count := map[byte]int{}
	left := 0
	for right, subStr := range s {
		count[byte(subStr)]++
		for count[byte(subStr)] > 1 { // 判断是否重复
			// 说明有重复
			// 恢复数据
			count[s[left]]--
			// 左端点右移动
			left += 1
		}
		ans = max(ans, right-left+1)
	}

	return ans
}
