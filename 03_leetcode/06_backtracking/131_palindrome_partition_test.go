package _6_backtracking

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 分割回文串 https://leetcode.cn/problems/palindrome-partitioning/?envType=study-plan-v2&envId=top-100-liked
func TestPartition(t *testing.T) {
	convey.Convey("分割回文串：想象两个字母之间是逗号", t, func() {
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

// 方法一：输入的视角（逗号选或不选）
func partition(s string) (ans [][]string) {
	n := len(s)
	path := []string{}

	// start 表示当前这段回文子串的开始位置
	// i 代表遍历时终点位置
	var dfs func(int, int)
	dfs = func(i, start int) {
		if i == n {
			ans = append(ans, append([]string(nil), path...)) // 复制 path
			return
		}

		if i < n-1 { // 不选 i 和 i+1 之间的逗号（i=n-1 时一定要选）
			dfs(i+1, start)
		}

		// 选 i 和 i+1 之间的逗号（把 s[i] 作为子串的最后一个字符）
		if isPalindrome(s, start, i) { // 回文情况下才选
			path = append(path, s[start:i+1])
			dfs(i+1, i+1)             // 下一个子串从 i+1 开始
			path = path[:len(path)-1] // 恢复现场
		}
	}

	dfs(0, 0)
	return
}

// 方法二： 答案的视角（枚举子串结束位置）
func partition2(s string) (ans [][]string) {
	var path []string
	n := len(s)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			// 与 78 子集对比：这里每个字母都在答案中，所以不是一递归就放在答案中，所以要长度相等
			ans = append(ans, append([]string(nil), path...)) // 复制 path
			return
		}
		for j := i; j < n; j++ { // 枚举子串的结束位置
			if isPalindrome(s, i, j) { // 回文的情况
				path = append(path, s[i:j+1]) // 这里注意是左闭右开区间
				dfs(j + 1)
				path = path[:len(path)-1] // 恢复现场
			}
		}
	}
	dfs(0)
	return
}

func isPalindrome(s string, left, right int) bool {
	// 使用相向指针
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
