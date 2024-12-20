package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 二叉树的右视图：https://leetcode.cn/problems/binary-tree-right-side-view/description/?envType=study-plan-v2&envId=top-100-liked
func TestRightSideView(t *testing.T) {
	convey.Convey("二叉树的右视图", t, func() {
		testCase := []struct {
			input  *TreeNode
			target []int
		}{
			{
				Ints2TreeNode([]int{1, 2, 3, 0, 5, 0, 4}),
				[]int{1, 3, 4},
			},
			{
				&TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 3,
					},
				},
				[]int{1, 3},
			},
		}

		for _, tst := range testCase {
			rsp := rightSideView(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

// 取每一层的最后一个节点
func rightSideView(root *TreeNode) (ans []int) {
	// 先递归右子树，再递归左子树（如果递归深度==答案的长度，说明可以记录）

	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == len(ans) {
			// 当递归深度==答案的长度，就记下来了
			ans = append(ans, node.Val)
		}
		depth++
		// 先右子树
		dfs(node.Right, depth)
		// 再左子树
		dfs(node.Left, depth)
	}
	dfs(root, 0)
	return
}
