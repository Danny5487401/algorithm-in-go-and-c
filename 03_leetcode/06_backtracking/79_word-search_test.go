package _6_backtracking

import (
	"github.com/smartystreets/goconvey/convey"
	"slices"
	"testing"
)

// 单词搜索 https://leetcode.cn/problems/word-search/description/?envType=study-plan-v2&envId=top-100-liked
func TestExist(t *testing.T) {
	convey.Convey("通过相邻的单元格内的字母构成", t, func() {
		testCase := []struct {
			input  [][]byte
			word   string
			target bool
		}{
			{

				[][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				}, "ABCCED", true,
			},

			{

				[][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				}, "SEE", true,
			},
			{

				[][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				}, "ABCB", false,
			}}

		for _, tst := range testCase {
			rsp := exist(tst.input, tst.word)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

type pair struct{ x, y int }

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 四方向

func exist(board [][]byte, word string) bool {
	cnt := map[byte]int{}
	for _, row := range board {
		for _, c := range row {
			cnt[c]++
		}
	}

	// 优化一: 如果 word 的某个字母的出现次数，比 board 中的这个字母的出现次数还要多，可以直接返回 false
	w := []byte(word)
	wordCnt := map[byte]int{}
	for _, c := range w {
		wordCnt[c]++
		if wordCnt[c] > cnt[c] {
			return false
		}
	}

	// 优化二: word 的第一个字母在 board 中出现了 x 次，word 的最后一个字母在 board 中出现了 y 次
	// 可以把 word 反转，相当于从 word 的最后一个字母开始搜索，这样更容易在一开始就满足 board[i][j] != word[k]
	if cnt[w[len(w)-1]] < cnt[w[0]] {
		slices.Reverse(w)
	}

	// 1 <= m, n <= 6
	m, n := len(board), len(board[0])
	var dfs func(i, j, k int) bool
	dfs = func(i, j int, k int) bool {
		if board[i][j] != word[k] { // 匹配失败
			return false
		}
		if k == len(word)-1 { // 匹配成功
			return true
		}
		// 避免重复访问同一个格子，覆盖标记
		board[i][j] = 1 // 标记访问过,word 仅由大小写英文字母组成
		for _, direction := range directions {
			x, y := i+direction.x, j+direction.y
			if 0 <= x && x < m && 0 <= y && y < n && dfs(x, y, k+1) {
				return true // 搜到了！
			}

		}
		board[i][j] = word[k] // 恢复现场
		return false          // 没搜到
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true // 搜到了！
			}
		}
	}
	return false
}
