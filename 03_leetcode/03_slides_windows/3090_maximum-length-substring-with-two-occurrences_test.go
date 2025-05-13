package _3_slides_windows

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 每个字符最多出现两次的最长子字符串: https://leetcode.cn/problems/maximum-length-substring-with-two-occurrences/description/
func TestMaximumLengthSubstring(t *testing.T) {
	convey.Convey(" 每个字符最多出现两次的最长子字符串", t, func() {
		testCase := []struct {
			input    string
			expected int
		}{
			{
				"bcbbbcba", 4,
			},
			{
				"aaaa", 2,
			},
		}

		for _, tst := range testCase {
			rsp := maximumLengthSubstring(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.expected)
		}
	})

}

func maximumLengthSubstring(s string) int {
	var numCount = map[byte]int{}

	var left int
	maxNum := 0
	for right, str := range s {
		numCount[byte(str)]++

		for numCount[byte(str)] > 2 {
			numCount[byte(s[left])]--
			left++
		}
		maxNum = max(maxNum, right-left+1)

	}
	return maxNum
}
