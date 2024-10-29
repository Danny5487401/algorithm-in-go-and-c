package _6_backtracking

import (
	"github.com/smartystreets/goconvey/convey"
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
		}

		for _, tst := range testCase {
			rsp := restoreIpAddresses(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

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
