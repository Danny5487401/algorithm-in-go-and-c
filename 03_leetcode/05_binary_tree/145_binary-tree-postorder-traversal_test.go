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
			rsp := postorderTraversal3(tst.input)
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

// 方式二: 迭代 https://leetcode.cn/problems/binary-tree-postorder-traversal/solutions/431066/er-cha-shu-de-hou-xu-bian-li-by-leetcode-solution/
func postorderTraversal3(root *TreeNode) []int {
	var stack []*TreeNode
	var prev *TreeNode
	var ans []int
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Right == nil || node.Right == prev {
			ans = append(ans, node.Val)
			prev = node
			node = nil
		} else { // 有右节点
			stack = append(stack, node)
			node = node.Right
		}

	}

	return ans

}
