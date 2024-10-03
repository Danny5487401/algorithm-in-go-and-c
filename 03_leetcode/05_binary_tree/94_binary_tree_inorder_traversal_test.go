package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 二叉树的中序遍历 https://leetcode.cn/problems/binary-tree-inorder-traversal/description/?envType=study-plan-v2&envId=top-100-liked
func TestInorderTraversal(t *testing.T) {
	convey.Convey("二叉树的中序遍历", t, func() {
		testCase := []struct {
			input  *TreeNode
			target []int
		}{
			{
				Ints2TreeNode([]int{1, NULL, 2, 3}),
				[]int{1, 3, 2},
			},
			{
				&TreeNode{
					Val: 1,
				},
				[]int{1},
			},
		}

		for _, tst := range testCase {
			rsp := inorderTraversal2(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func inorderTraversal(root *TreeNode) []int {
	var ans = make([]int, 0)
	if root == nil {
		return ans
	}
	// 先遍历左子树
	left := inorderTraversal(root.Left)
	ans = append(ans, left...)

	// 再中间
	ans = append(ans, root.Val)

	// 最后右子树
	right := inorderTraversal(root.Right)
	ans = append(ans, right...)
	return ans
}

func inorderTraversal2(root *TreeNode) []int {
	var ans = make([]int, 0)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		ans = append(ans, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return ans
}
