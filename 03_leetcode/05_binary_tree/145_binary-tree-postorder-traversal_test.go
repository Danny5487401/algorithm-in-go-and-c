package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 二叉树的后序遍历 https://leetcode.cn/problems/binary-tree-postorder-traversal/description/
func TestPostorderTraversal(t *testing.T) {
	convey.Convey("二叉树的后序遍历", t, func() {
		testCase := []struct {
			input  *TreeNode
			target []int
		}{
			{
				Ints2TreeNode([]int{1, NULL, 2, 3}),
				[]int{3, 2, 1},
			},
			{
				Ints2TreeNode([]int{1, 2, 3, 4, 5, NULL, 8, NULL, NULL, 6, 7, 9}),
				[]int{4, 6, 7, 5, 2, 9, 8, 3, 1},
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
			rsp := postorderTraversal(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func postorderTraversal(root *TreeNode) []int {
	/*
		- 后序遍历: [ [左子树的前序遍历结果], [右子树的前序遍历结果],根节点 ]
	*/
	var ans []int

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		dfs(root.Right)
		ans = append(ans, root.Val)
	}
	dfs(root)
	return ans
}
