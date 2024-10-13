package _5_binary_tree

// 层数最深叶子节点的和 https://leetcode.cn/problems/deepest-leaves-sum/
func deepestLeavesSum(root *TreeNode) (sum int) {
	if root == nil {
		return 0
	}
	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		sum = 0
		value := make([]*TreeNode, len(queue))
		for _ = range value {
			cur := queue[0]
			queue = queue[1:]
			sum += cur.Val
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}

	}
	return
}
