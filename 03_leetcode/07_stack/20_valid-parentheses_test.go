package _7_stack

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 有效的括号 https://leetcode.cn/problems/valid-parentheses/description/?envType=study-plan-v2&envId=top-100-liked
func TestIsValid(t *testing.T) {
	convey.Convey("有效的括号", t, func() {
		testCase := []struct {
			input    string
			expected bool
		}{
			{
				"()", true,
			},
			{
				"()[]{}", true,
			},
			{
				"(]", false,
			},
			{
				"([])", true,
			},
		}

		for _, tst := range testCase {
			rsp := isValid(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.expected)
		}
	})

}

func isValid(s string) bool {
	if len(s)%2 != 0 { // s 长度必须是偶数
		return false
	}
	mp := map[rune]rune{')': '(', ']': '[', '}': '{'} // s 仅由括号 '()[]{}' 组成
	var st []rune
	for _, c := range s {
		if mp[c] == 0 { // c 是左括号
			st = append(st, c) // 入栈
		} else { // c 是右括号
			if len(st) == 0 || st[len(st)-1] != mp[c] {
				return false // 没有左括号，或者左括号类型不对
			}
			st = st[:len(st)-1] // 出栈
		}
	}
	return len(st) == 0 // 所有左括号必须匹配完毕
}
