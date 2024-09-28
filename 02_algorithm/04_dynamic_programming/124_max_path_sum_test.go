package _4_dynamic_programming

import (
	"math"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

// 二叉树中的最大路径和 https://leetcode.cn/problems/binary-tree-maximum-path-sum/description/

func TestMaxPathSum(t *testing.T) {
	convey.Convey("二叉树中的最大路径和", t, func() {
		testCase := []struct {
			input  *TreeNode
			target int
		}{
			{
				Ints2TreeNode([]int{1, 2, 3}),
				6,
			},
			{
				Ints2TreeNode([]int{-10, 9, 20, NULL, NULL, 15, 7}),
				42,
			},
		}

		for _, tst := range testCase {
			rsp := maxPathSum(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func maxPathSum(root *TreeNode) int {

	var dfs func(root *TreeNode) int

	var ans = math.MinInt
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0 // 没有节点，和为 0
		}
		leftSum := dfs(node.Left)
		rightSum := dfs(node.Right)
		// 当前拐弯节点的最大路径和=左子树最大链和 + 右子树最大链和 + 当前节点值
		ans = max(ans, leftSum+rightSum+node.Val)

		// 返回父节点的 max(左子树最大链和 , 右子树最大链和 ) + 当前节点值
		temRsp := max(leftSum, rightSum) + node.Val
		return max(temRsp, 0) // 如果小于0,表示不选

	}
	dfs(root)
	return ans
}
