package _5_binary_tree

// 对称二叉树 https://leetcode.cn/problems/symmetric-tree/?envType=study-plan-v2&envId=top-100-liked

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	// 平分根节点，就与 100. 相同的树 相似：https://leetcode.cn/problems/same-tree/
	return isSymmetricTree(root.Left, root.Right)
}

func isSymmetricTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	// 判断依据 左子树=右子树
	return p.Val == q.Val && isSymmetricTree(p.Left, q.Right) && isSymmetricTree(p.Right, q.Left)
}
