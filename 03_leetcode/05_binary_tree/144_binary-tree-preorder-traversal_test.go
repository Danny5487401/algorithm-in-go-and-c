package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 二叉树的前序遍历 https://leetcode.cn/problems/binary-tree-preorder-traversal/
func TestPreorderTraversal(t *testing.T) {
	convey.Convey("二叉树的前序遍历", t, func() {
		testCase := []struct {
			input  *TreeNode
			target []int
		}{
			{
				Ints2TreeNode([]int{1, NULL, 2, 3}),
				[]int{1, 2, 3},
			},
			{
				Ints2TreeNode([]int{1, 2, 3, 4, 5, NULL, 8, NULL, NULL, 6, 7, 9}),
				[]int{1, 2, 4, 5, 6, 7, 3, 8, 9},
			},
			{
				Ints2TreeNode([]int{}),
				[]int(nil),
			},
			{
				Ints2TreeNode([]int{1}),
				[]int{1},
			},
		}

		for _, tst := range testCase {
			rsp := preorderTraversal(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func preorderTraversal(root *TreeNode) []int {
	/*
		- 前序遍历: [ 根节点, [左子树的前序遍历结果], [右子树的前序遍历结果] ]
	*/
	var ans []int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		ans = append(ans, root.Val)
		dfs(root.Left)
		dfs(root.Right)

	}
	dfs(root)
	return ans

}
