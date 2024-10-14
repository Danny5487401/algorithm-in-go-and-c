package _6_backtracking

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 路径总和 II https://leetcode.cn/problems/path-sum-ii/
func TestHasPathSumII(t *testing.T) {
	convey.Convey("路径总和II:找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径", t, func() {
		testCase := []struct {
			input  *TreeNode
			sum    int
			target [][]int
		}{
			{
				Ints2TreeNode([]int{5, 4, 8, 11, NULL, 13, 4, 7, 2, NULL, NULL, 5, 1}),
				22,
				[][]int{
					{5, 4, 11, 2},
					{5, 8, 4, 5},
				},
			},
			{
				Ints2TreeNode([]int{1, 2, 3}),
				5,
				[][]int(nil),
			},
			{
				Ints2TreeNode([]int{1, 2}),
				0,
				[][]int(nil),
			},
		}

		for _, tst := range testCase {
			rsp := pathSumII(tst.input, tst.sum)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func pathSumII(root *TreeNode, targetSum int) [][]int {
	var dfs func(*TreeNode, int)

	var ans [][]int
	var path []int
	dfs = func(node *TreeNode, left int) {
		if node == nil {
			return
		}
		//if left <= 0 {
		//	return
		//}
		path = append(path, node.Val)
		left -= node.Val
		defer func() { path = path[:len(path)-1] }() // 恢复现场
		if node.Left == nil && node.Right == nil && left == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		dfs(node.Left, left)
		dfs(node.Right, left)

	}
	dfs(root, targetSum)
	return ans
}
