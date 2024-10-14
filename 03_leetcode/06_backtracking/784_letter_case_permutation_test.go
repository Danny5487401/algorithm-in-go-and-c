package _6_backtracking

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
	"unicode"
)

// 字母大小写全排列 https://leetcode.cn/problems/letter-case-permutation/
func TestLetterCasePermutation(t *testing.T) {
	convey.Convey("字母大小写全排列", t, func() {
		testCase := []struct {
			input  string
			target []string
		}{
			{

				"a1b2", []string{"a1b2", "a1B2", "A1b2", "A1B2"},
			},
			{

				"3z4", []string{"3z4", "3Z4"},
			},
		}

		for _, tst := range testCase {
			rsp := letterCasePermutation(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func letterCasePermutation(s string) []string {
	/*
		每个字符的大小写形式刚好差了 32，因此在大小写装换时可以用 c32 来进行转换和恢复
	*/
	length := len(s)
	path := make([]byte, length)

	var ans []string
	var dfs func(int)
	dfs = func(i int) {
		if i == length {
			ans = append(ans, string(path))
			return
		}
		path[i] = s[i]
		dfs(i + 1)

		if unicode.IsLetter(rune(s[i])) {
			path[i] ^= 32
			dfs(i + 1)
		}

	}
	dfs(0)
	return ans
}
