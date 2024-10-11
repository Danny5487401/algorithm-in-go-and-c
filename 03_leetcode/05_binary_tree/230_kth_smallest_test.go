package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 二叉搜索树中第K小的元素 https://leetcode.cn/problems/kth-smallest-element-in-a-bst/description/
func TestKthSmallest(t *testing.T) {
	convey.Convey("二叉搜索树中第K小的元素", t, func() {
		testCase := []struct {
			input  *TreeNode
			k      int
			target int
		}{
			{
				Ints2TreeNode([]int{3, 1, 4, NULL, 2}),
				1,
				1,
			},
			{
				Ints2TreeNode([]int{5, 3, 6, 2, 4, NULL, NULL, 1}),
				3,
				3,
			},
		}

		for _, tst := range testCase {
			rsp := kthSmallest(tst.input, tst.k)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{root}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
	return 0
}
