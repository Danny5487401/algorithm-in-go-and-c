package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 反转字符串 https://leetcode.cn/problems/reverse-string/description/
func TestReverseString(t *testing.T) {
	convey.Convey("反转字符串", t, func() {
		testCase := []struct {
			input    []byte
			expected []byte
		}{
			{

				[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'},
			},
			{
				[]byte{'H', 'a', 'n', 'n', 'a', 'h'}, []byte{'h', 'a', 'n', 'n', 'a', 'H'},
			},
		}

		for _, tst := range testCase {
			reverseString(tst.input)
			convey.So(tst.input, convey.ShouldResemble, tst.expected)
		}
	})
}
func reverseString(s []byte) {
	right := len(s) - 1
	for left := 0; left < right; left++ {
		s[left], s[right] = s[right], s[left]
		right--
	}
	return

}
