package _5_binary_tree

// 根据前序和后序遍历构造二叉树 https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-postorder-traversal/
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	/*
			- 前序遍历: [ 根节点, [左子树的前序遍历结果], [右子树的前序遍历结果] ]
			- 后序遍历: [ [左子树的中序遍历结果], [右子树的中序遍历结果],根节点 ]

		preLeft, preRight, postLeft, postRight

		可以对 preorder 和 postorder 进行分治处理，即将 preorder 划分为根节点、左子树节点和右子树节点三个部分，postorder 也划分为左子树节点、右子树节点和根节点三个部分。

	*/
	return nil

}
