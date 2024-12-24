package _7_stack

import (
	"github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
	"unicode"
)

// 字符串解码 https://leetcode.cn/problems/decode-string/description/?envType=study-plan-v2&envId=top-100-liked
func TestDecodeString(t *testing.T) {
	convey.Convey("字符串解码:编码规则 k[encoded_string]: 表示其中方括号内部的 encoded_string 正好重复 k 次", t, func() {
		testCase := []struct {
			nums     string
			expected string
		}{
			{
				"3[a]2[bc]", "aaabcbc",
			},
			{
				"3[a2[c]]", "accaccacc",
			},
			{
				"2[abc]3[cd]ef", "abcabccdcdcdef",
			},
		}

		for _, tst := range testCase {
			rsp := decodeString(tst.nums)
			convey.So(rsp, convey.ShouldEqual, tst.expected)
		}
	})

}

// 每遇到一个中括号为新的一层（layer）
type layer struct {
	res string
	mul int // 重复次数
}

// 用结构体单栈
func decodeString(s string) string {
	// 括号嵌套的情况，比如 2[a2[bc]]
	st := []layer{}
	curr := layer{"", 0}
	for _, c := range s {
		if unicode.IsDigit(c) {
			// 计算倍数
			curr.mul = curr.mul*10 + int(c-'0')
		} else if c == '[' {
			// 入栈
			st = append(st, curr)
			curr = layer{"", 0}
		} else if c == ']' {
			// 出栈
			prev := st[len(st)-1]
			st = st[:len(st)-1]
			curr.res = prev.res + strings.Repeat(curr.res, prev.mul)
		} else { // 字母a-z
			// 内部
			curr.res += string(c)
		}
	}
	return curr.res
}

// 方式二:双栈是同进同出
