package _3_slides_windows

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 定长子串中元音的最大数目 https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/description/
func TestMaxVowels(t *testing.T) {
	convey.Convey("请返回字符串 s 中长度为 k 的单个子字符串中可能包含的最大元音字母数", t, func() {
		testCase := []struct {
			nums     string
			k        int
			expected int
		}{
			{
				// iii
				"abciiidef", 3, 3,
			},
			{
				"aeiou", 2, 2,
			},
			{
				"leetcode", 3, 2,
			},
			{
				"rhythms", 4, 0,
			},
			{
				"tryhard", 4, 1,
			},
		}

		for _, tst := range testCase {
			rsp := maxVowels(tst.nums, tst.k)
			convey.So(rsp, convey.ShouldEqual, tst.expected)
		}
	})

}

func maxVowels(s string, k int) int {
	vowel := map[int32]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}
	var ans int
	maxVow := 0
	for i, in := range s {
		// 1. 进入窗口
		// 只需要考虑移除（离开窗口）的字母 a 是不是元音，以及添加（进入窗口）的字母 i 是不是元音
		if _, ok := vowel[in]; ok {
			maxVow++
		}
		if i < k-1 { // 窗口大小不足 k
			continue
		}
		// 2. 更新答案
		ans = max(ans, maxVow)
		// 3. 离开窗口
		out := s[i-k+1]
		if _, ok := vowel[int32(out)]; ok {
			maxVow--
		}
	}
	return ans

}
