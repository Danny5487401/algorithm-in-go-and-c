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
			rsp := inorderTraversal3(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

// 方式一:递归
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

// 方式二: 迭代
func inorderTraversal3(root *TreeNode) (res []int) {
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}

// 方式三: Morris 中序遍历
