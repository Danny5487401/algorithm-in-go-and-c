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
			rsp := preorderTraversal3(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

// 递归的时候隐式地维护了一个栈，而我们在迭代的时候需要显式地将这个栈模拟出来，

// 递归
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

func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ans := make([]int, 0)

	ans = append(ans, root.Val)
	ans = append(ans, preorderTraversal2(root.Left)...)
	ans = append(ans, preorderTraversal2(root.Right)...)
	return ans
}

// 迭代 https://leetcode.cn/problems/binary-tree-preorder-traversal/solutions/461821/er-cha-shu-de-qian-xu-bian-li-by-leetcode-solution/
func preorderTraversal3(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var stack []*TreeNode
	node := root
	ans := make([]int, 0)
	for len(stack) > 0 || node != nil {
		for node != nil {
			ans = append(ans, node.Val) // 先节点本身
			stack = append(stack, node) // 一直入栈左节点
			node = node.Left            // 左节点
		}
		// 出栈
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]

	}
	return ans
}
