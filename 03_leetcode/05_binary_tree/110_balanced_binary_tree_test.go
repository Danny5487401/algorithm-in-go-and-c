package _5_binary_tree

import "math"

// 平衡二叉树 https://leetcode.cn/problems/balanced-binary-tree/description/
func isBalanced(root *TreeNode) bool {
	// 高度为正值，这里假设高度是 -1 ，提前返回
	return maxBalancedDepth(root) != -1
}

func maxBalancedDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 递归 DFS
	// 平衡的要求是高度差不大于1

	// 分别左子树和右子树的最大深度 l 和 r
	leftDepth := maxBalancedDepth(root.Left)
	if leftDepth == -1 {
		return -1
	}
	rightDepth := maxBalancedDepth(root.Right)
	if rightDepth == -1 || math.Abs(float64(rightDepth)-float64(leftDepth)) > 1 {
		return -1
	}
	return max(leftDepth, rightDepth) + 1

}
