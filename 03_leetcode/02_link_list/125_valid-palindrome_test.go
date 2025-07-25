package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
	"unicode"
)

// 验证回文串 https://leetcode.cn/problems/valid-palindrome/
func TestIsPalindromeCheck(t *testing.T) {
	convey.Convey("如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 ", t, func() {
		testCase := []struct {
			input    string
			expected bool
		}{
			{

				"A man, a plan, a canal: Panama", true,
			},
			{
				"race a car", false,
			},
			{
				" ", true,
			},
		}

		for _, tst := range testCase {
			rsp := isPalindromeCheck(tst.input)
			convey.So(rsp, convey.ShouldResemble, tst.expected)
		}
	})
}
func isPalindromeCheck(s string) bool {
	s = strings.ToLower(s)
	newStr := []byte{}
	for _, v := range s {
		if unicode.IsLetter(v) || unicode.IsNumber(v) {
			newStr = append(newStr, byte(v))
		}
	}
	right := len(newStr) - 1
	for left := 0; left < right; left++ {
		if newStr[left] != newStr[right] {
			return false
		}
		right--
	}
	return true

}
