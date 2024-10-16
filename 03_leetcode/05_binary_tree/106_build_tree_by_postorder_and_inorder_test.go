package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 从中序与后序遍历序列构造二叉树 https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
func TestBuildTreeByInorderAndPostOrder(t *testing.T) {
	convey.Convey("从中序与后序遍历序列构造二叉树", t, func() {
		testCase := []struct {
			inorder   []int
			postorder []int
			target    *TreeNode
		}{
			{
				[]int{9, 3, 15, 20, 7},
				[]int{9, 15, 7, 20, 3},
				Ints2TreeNode([]int{3, 9, 20, NULL, NULL, 15, 7}),
			},
			{
				[]int{-1},
				[]int{-1},
				Ints2TreeNode([]int{-1}),
			},
		}

		for _, tst := range testCase {
			rsp := buildTreeByInorderAndPostOrder(tst.inorder, tst.postorder)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func buildTreeByInorderAndPostOrder(inorder []int, postorder []int) *TreeNode {
	/*
		- 中序遍历: [ [左子树的中序遍历结果], 根节点, [右子树的中序遍历结果] ]
		- 后序遍历: [ [左子树的中序遍历结果], [右子树的中序遍历结果],根节点 ]
	*/
	if len(postorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}
	}
	root.Left = buildTreeByInorderAndPostOrder(inorder[:i], postorder[:i])
	root.Right = buildTreeByInorderAndPostOrder(inorder[i+1:], postorder[i:len(postorder)-1])

	return root
}
