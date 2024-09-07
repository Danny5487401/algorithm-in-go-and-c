package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 路径总和 https://leetcode.cn/problems/path-sum/description/
func TestHasPathSum(t *testing.T) {
	convey.Convey("路径总和:是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum", t, func() {
		testCase := []struct {
			input  *TreeNode
			sum    int
			target bool
		}{
			{
				Ints2TreeNode([]int{5, 4, 8, 11, NULL, 13, 4, 7, 2, NULL, NULL, NULL, 1}),
				22,
				true,
			},
			{
				Ints2TreeNode([]int{1, 2, 3}),
				5,
				false,
			},
			{
				Ints2TreeNode([]int{}),
				0,
				false,
			},
		}

		for _, tst := range testCase {
			rsp := hasPathSum(tst.input, tst.sum)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func hasPathSum(root *TreeNode, targetSum int) bool {

	if root == nil {
		return false // Since the tree is empty, there are no root-to-leaf paths
	}

	targetSum -= root.Val

	if root.Left == nil && root.Right == nil { // 叶子节点
		return targetSum == 0
	}
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}
