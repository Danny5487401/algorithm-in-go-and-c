package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 从前序与中序遍历序列构造二叉树 https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
func TestBuildTree(t *testing.T) {
	convey.Convey("从前序与中序遍历序列构造二叉树", t, func() {
		testCase := []struct {
			preorder []int
			inorder  []int
			target   *TreeNode
		}{
			{
				[]int{3, 9, 20, 15, 7},
				[]int{9, 3, 15, 20, 7},
				Ints2TreeNode([]int{3, 9, 20, NULL, NULL, 15, 7}),
			},
			{
				[]int{-1},
				[]int{-1},
				Ints2TreeNode([]int{-1}),
			},
		}

		for _, tst := range testCase {
			rsp := buildTree(tst.preorder, tst.inorder)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	// 根节点
	root := &TreeNode{Val: preorder[0]}
	i := 0
	for ; i < len(inorder); i++ {
		// 查找跟节点,划分左右边界
		if inorder[i] == root.Val {
			break
		}
	}
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}
