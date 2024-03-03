package _5_binary_tree

// 二叉树的右视图：https://leetcode.cn/problems/binary-tree-right-side-view/description/?envType=study-plan-v2&envId=top-100-liked
// 取每一层的最后一个节点

func rightSideView(root *TreeNode) (ans []int) {
	// 先递归右子树，再递归左子树（如果递归深度==答案的长度，说明可以记录）

	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == len(ans) {
			ans = append(ans, node.Val)
		}
		// 先右子树
		dfs(node.Right, depth+1)
		dfs(node.Left, depth+1)
	}
	dfs(root, 0)
	return
}
