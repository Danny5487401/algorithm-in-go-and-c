package _3_slides_windows

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 找到最长的半重复子字符串 https://leetcode.cn/problems/find-the-longest-semi-repetitive-substring/

func TestLongestSemiRepetitiveSubstring(t *testing.T) {
	convey.Convey("找到最长的半重复子字符串:一个字符串 t 中至多有一对相邻字符是相等的，那么称这个字符串 t 是 半重复的 ", t, func() {
		testCase := []struct {
			input    string
			expected int
		}{
			{
				"52233",
				4,
			},
			{
				"5494",
				4,
			},
			{
				"1111111",
				2,
			},
		}

		for _, tst := range testCase {
			rsp := longestSemiRepetitiveSubstring(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.expected)
		}
	})

}

func longestSemiRepetitiveSubstring(s string) int {
	var count int // 统计相邻相同的情况出现了多少次

	var left int

	var ans = 1
	for right := 1; right < len(s); right++ {

		if s[right-1] == s[right] {
			count++

			if count > 1 {
				left++

				for s[left] != s[left-1] { // 移动左指针 left 直到 s[left]=s[left−1]，此时将一对相同的字符移到窗口之外
					// 比如 1101场景, 100211 场景
					left++
				}

			}

		}

		ans = max(ans, right-left+1)
	}
	return ans
}
