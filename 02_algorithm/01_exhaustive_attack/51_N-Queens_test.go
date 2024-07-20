package _1_exhaustive_attack

import (
	"github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

// N 皇后 https://leetcode.cn/problems/n-queens/description/

func TestPermute(t *testing.T) {
	convey.Convey("每行每列只能有一个皇后, 不能同一斜线", t, func() {
		testCase := []struct {
			input  int
			target [][]string
		}{
			{
				4, [][]string{
					{
						".Q..", "...Q", "Q...", "..Q.",
					},
					{
						"..Q.", "Q...", "...Q", ".Q..",
					},
				},
			},
			{

				1, [][]string{
					{
						"Q",
					},
				},
			},
		}

		for _, tst := range testCase {
			rsp := solveNQueens(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func solveNQueens(n int) [][]string {
	var ans = make([][]string, 0)

	// 记录列 column 是否使用过
	onPathColumn := make([]bool, n)

	// 斜线的两个方向 column + row  和  column - row
	diag1 := make([]bool, n*2-1)
	diag2 := make([]bool, n*2-1)

	column := make([]int, n)

	// row 代表当前要枚举的行号，column代表剩余可以枚举的列号
	var dfs func(r int)
	dfs = func(row int) {
		if row == n { // 所有行号都遍历完毕
			board := make([]string, n)
			for i, c := range column {
				// 遍历列，写数据
				board[i] = strings.Repeat(".", c) + "Q" + strings.Repeat(".", n-c-1)
			}
			ans = append(ans, board)
			return
		}
		for c, on := range onPathColumn {
			rSubC := row - c + (n - 1)
			rAddC := row + c
			if !on && !diag1[rSubC] && !diag2[rAddC] { // 选择没有列举过列
				column[row] = c
				onPathColumn[c], diag1[rSubC], diag2[rAddC] = true, true, true

				dfs(row + 1)
				onPathColumn[c], diag1[rSubC], diag2[rAddC] = false, false, false
			}
		}

	}
	dfs(0)
	return ans
}
