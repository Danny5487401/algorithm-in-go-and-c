package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

// 平衡二叉树 https://leetcode.cn/problems/balanced-binary-tree/description/
func TestIsBalanced(t *testing.T) {
	convey.Convey("平衡二叉树", t, func() {
		testCase := []struct {
			input  *TreeNode
			target bool
		}{
			{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
							Left: &TreeNode{
								Val: 4,
							},
							Right: &TreeNode{
								Val: 4,
							},
						},
						Right: &TreeNode{
							Val: 3,
						},
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
				false,
			},
			{
				nil,
				true,
			},
		}

		for _, tst := range testCase {
			rsp := isBalanced(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func isBalanced(root *TreeNode) bool {
	// 高度为正值，这里假设高度是 -1 ，如果不等于-1可以提前返回
	return maxBalancedDepth(root) != -1
}

func maxBalancedDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 递归 DFS

	// 分别左子树和右子树的最大深度 l 和 r
	leftDepth := maxBalancedDepth(root.Left)
	if leftDepth == -1 {
		// 发现不平衡
		return -1
	}
	rightDepth := maxBalancedDepth(root.Right)
	if rightDepth == -1 || math.Abs(float64(rightDepth)-float64(leftDepth)) > 1 {
		// 平衡的要求是高度差不大于1
		return -1
	}
	return max(leftDepth, rightDepth) + 1

}
