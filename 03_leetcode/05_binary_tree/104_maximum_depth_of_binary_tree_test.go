package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 二叉树的最大深度 https://leetcode.cn/problems/maximum-depth-of-binary-tree/?envType=study-plan-v2&envId=top-100-liked

func TestMaxDepth(t *testing.T) {
	convey.Convey("二叉树的最大深度", t, func() {
		testCase := []struct {
			input  *TreeNode
			target int
		}{
			{
				&TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val: 20,
						Left: &TreeNode{
							Val: 15,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
				3,
			},
		}

		for _, tst := range testCase {
			rsp := maxDepth(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 递归 DFS max(leftDepth, rightDepth) + 1
	// 左子树和右子树的最大深度 l 和 r
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	return max(leftDepth, rightDepth) + 1

}

// 方法二：用全局变量保存最大值,传递节点个数
func maxDepth2(root *TreeNode) int {
	var ans = 0
	var dfs func(node *TreeNode, count int)
	dfs = func(node *TreeNode, count int) {
		if node == nil {
			return
		}
		// 每次到一个节点 +1
		count++
		ans = max(ans, count)

		// 递归左右子树
		dfs(node.Left, count)
		dfs(node.Right, count)
	}
	dfs(root, 0)
	return ans

}
