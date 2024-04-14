package _4_dynamic_programming

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 二叉树的直径 https://leetcode.cn/problems/diameter-of-binary-tree/description/

func TestDiameterOfBinaryTree(t *testing.T) {
	convey.Convey("二叉树的直径", t, func() {
		testCase := []struct {
			input  *TreeNode
			target int
		}{
			{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				3,
			},
			{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
				},
				1,
			},
		}

		for _, tst := range testCase {
			rsp := diameterOfBinaryTree(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func diameterOfBinaryTree(root *TreeNode) int {

	var dfs func(node *TreeNode) int
	var ans int
	dfs = func(node *TreeNode) int {
		if node == nil {
			// 当叶子节点时
			return -1
		}
		leftLength := dfs(node.Left)
		rightLength := dfs(node.Right)
		// 当前拐弯的直接长度 = 左子路的最长链 + 右子路的最长链 + 2
		ans = max(ans, leftLength+rightLength+2)

		// 返回父节点节点 max(leftLength, rightLength) + 1
		return max(leftLength, rightLength) + 1
	}
	dfs(root)
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
