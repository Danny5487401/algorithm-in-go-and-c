package _5_binary_tree

// 相同的树 https://leetcode.cn/problems/same-tree/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// 边界条件
	if p == nil || q == nil {
		return p == q
	}
	// 判断依据是左子树=左子树
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
