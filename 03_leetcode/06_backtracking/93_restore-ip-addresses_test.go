package _6_backtracking

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// 复原 IP 地址 https://leetcode.cn/problems/restore-ip-addresses/description/
func TestRestoreIpAddresses(t *testing.T) {
	convey.Convey("复原 IP 地址", t, func() {
		testCase := []struct {
			input  string
			target []string
		}{
			{
				"25525511135",
				[]string{
					"255.255.11.135",
					"255.255.111.35",
				},
			},

			{
				"0000",
				[]string{
					"0.0.0.0",
				},
			},

			{
				"101023",
				[]string{
					"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3",
				},
			},
		}

		for _, tst := range testCase {
			rsp := restoreIpAddresses2(tst.input)
			sort.Strings(rsp)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

// 对比 131 分割回文串没有分割段数的要求，而本题要求分成恰好 4 段
/*
在 131 题的基础上，我们需要额外知道：

	当前在第几段。
	当前分割的子串，对应的数值是多少，便于我们判断子串是否合法。
定义 dfs(i,j,ipVal)，表示：

	剩余字符从 s[i] 到 s[n−1]。
	现在在第 j 段（j 从 0 开始）。
	第 j 段的数值目前为 ipVal。

递归过程中，首先把 s[i] 加到当前这一段的末尾，即更新 ipVal 为 ipVal⋅10+int(s[i])。例如在 12 的末尾添加 3，数值更新为 12⋅10+3=123。

分类讨论：
* 不分割，前提是不能有前导零，即此时 ipVal>0。往下递归到 dfs(i+1,j,ipVal)。注意，如果有前导零的话，会在前导零那个字符处发现 ipVal=0，不会往下递归。
* 分割，s[i] 作为当前这段子串的右端点。把 j 加一，ipVal 重置为 0。往下递归到 dfs(i+1,j+1,0)。

递归边界：
* i=n 时，s 分割完毕，如果此时 j=4，把分割结果加入答案。
* 否则，如果 j=4，由于此时已经分出 4 段，不能再分割，所以不再递归。
递归入口：dfs(0,0,0)。

*/

func restoreIpAddresses2(s string) (ans []string) {
	length := len(s)
	path := [4]int{}
	var dfs func(int, int, int)
	dfs = func(i int, j int, ipVal int) { // i 为字符串位置,j 为第几段,ipVal 第j 段数值
		if i == length {
			if j == 4 {
				a, b, c := path[0], path[1], path[2]
				ans = append(ans, fmt.Sprintf("%s.%s.%s.%s", s[:a], s[a:b], s[b:c], s[c:]))
				return
			}
			return
		}
		if j == 4 { // j=4 的时候必须分割完毕，不能有剩余字符
			return
		}
		// 手动把字符串转成整数，这样字符串转整数是严格 O(1) 的
		ipVal = ipVal*10 + int(s[i]-'0')
		if ipVal > 255 { // 不合法
			return
		}
		// 不分割，不以 s[i] 为这一段的结尾
		if ipVal > 0 { // 无前导零
			dfs(i+1, j, ipVal)
		}

		// 分割，以 s[i] 为这一段的结尾
		path[j] = i + 1 // 记录下一段的开始位置
		dfs(i+1, j+1, 0)

	}
	dfs(0, 0, 0)
	return
}

func restoreIpAddresses(s string) (ans []string) {
	n := len(s)
	var dfs func(subRes []string, start int)
	dfs = func(subRes []string, start int) {
		if len(subRes) == 4 && start == n { // 凑够了4段，且已遍历完s
			ans = append(ans, strings.Join(subRes, "."))
			return
		}
		// 剪枝优化 1
		if len(subRes) == 4 && start < n { // 凑够了4段，还未遍历完s，不符题意（必须用完所有字符的需求）
			return
		}
		for length := 1; length <= 3; length++ { // IP 范围为 0-255
			// 剪枝优化 2
			if start+length-1 >= n { // 若加上此段则超出s的长度，所以不能加
				return
			}
			numStr := s[start : start+length] // 截出来了此段

			// 剪枝优化 3
			if numStr[0] == '0' && length > 1 { // 此段可能为'0', 但不能为'0x' 或 '0xx'
				return
			}
			// 剪枝优化 4
			if num, _ := strconv.Atoi(numStr); num > 255 { // 值域需介于[0,255]
				return
			}

			subRes = append(subRes, numStr) // 下一轮选择
			dfs(subRes, start+length)
			subRes = subRes[:len(subRes)-1] // 回溯：撤回刚才的选择
		}
	}
	dfs([]string{}, 0)
	return
}
