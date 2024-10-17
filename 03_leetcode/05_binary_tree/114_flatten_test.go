package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 二叉树展开为链表 https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/description/?envType=study-plan-v2&envId=top-100-liked
func TestFlatten(t *testing.T) {
	convey.Convey(" 二叉树展开为链表", t, func() {
		testCase := []struct {
			input *TreeNode
		}{
			{
				Ints2TreeNode([]int{1, 2, 5, 3, 4, NULL, 6}),
			},
			{
				Ints2TreeNode([]int{}),
			},

			{
				Ints2TreeNode([]int{0}),
			},
		}

		for _, tst := range testCase {
			flatten(tst.input)
		}
	})

}

// 方式一:递归
func flatten(root *TreeNode) {

	var ans = make([]*TreeNode, 0)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		ans = append(ans, node)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	for i := 1; i < len(ans); i++ {
		prev, curr := ans[i-1], ans[i]
		prev.Left, prev.Right = nil, curr
	}

}
