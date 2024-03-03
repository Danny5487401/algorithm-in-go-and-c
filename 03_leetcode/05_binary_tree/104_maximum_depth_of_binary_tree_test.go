package _5_binary_tree

// 二叉树的最大深度 https://leetcode.cn/problems/maximum-depth-of-binary-tree/?envType=study-plan-v2&envId=top-100-liked

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 递归 DFS
	// 左子树和右子树的最大深度 l 和 r
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	return max(leftDepth, rightDepth) + 1

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
