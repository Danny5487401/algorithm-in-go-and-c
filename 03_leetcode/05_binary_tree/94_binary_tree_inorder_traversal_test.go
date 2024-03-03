package _5_binary_tree

// 二叉树的中序遍历 https://leetcode.cn/problems/binary-tree-inorder-traversal/description/?envType=study-plan-v2&envId=top-100-liked

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
