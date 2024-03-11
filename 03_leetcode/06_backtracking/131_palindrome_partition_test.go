package _6_backtracking

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 分割回文串  https://leetcode.cn/problems/palindrome-partitioning/?envType=study-plan-v2&envId=top-100-liked

func TestPartition(t *testing.T) {
	convey.Convey("分割回文串", t, func() {
		testCase := []struct {
			input  string
			target [][]string
		}{
			{

				"aab", [][]string{{"a", "a", "b"}, {"aa", "b"}},
			},
			{
				"a", [][]string{{"a"}},
			},
		}

		for _, tst := range testCase {
			rsp := partition(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

// 方法二： 答案的视角（枚举子串结束位置）
func partition(s string) (ans [][]string) {
	path := []string{}
	n := len(s)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string(nil), path...)) // 复制 path
			return
		}
		for j := i; j < n; j++ { // 枚举子串的结束位置
			if isPalindrome(s, i, j) {
				path = append(path, s[i:j+1])
				dfs(j + 1)
				path = path[:len(path)-1] // 恢复现场
			}
		}
	}
	dfs(0)
	return
}

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
